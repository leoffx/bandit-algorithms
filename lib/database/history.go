package database

import (
	"fmt"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
)

type Round struct {
	Probability float64
	Reward      float64
	WasChosen   bool
}

type Stats struct {
	Count                 int
	AvgRewardWhenUsed     float64
	AvgRewardWhenEligible float64
	MovingAvg             float64
	TimeSinceLast         int
	Rounds                []*Round
}

func (s Stats) String() string {
	return fmt.Sprintf("Count: %d, AvgRewardWhenUsed: %f, TimeSinceLast: %d", s.Count, s.AvgRewardWhenUsed, s.TimeSinceLast)
}

type History struct {
	ArmToStats map[*bandit.Arm]*Stats
}

func NewHistory(arms []*bandit.Arm) *History {
	armToStats := make(map[*bandit.Arm]*Stats)
	for _, arm := range arms {
		armToStats[arm] = &Stats{
			Rounds: []*Round{},
		}
	}
	return &History{
		ArmToStats: armToStats,
	}
}

func (h *History) Update(chosenArm *bandit.Arm, armToProbability map[*bandit.Arm]float64, reward float64) {
	// Order is important here
	h.updateRounds(chosenArm, armToProbability, reward)
	h.updateAvgRewards(armToProbability, reward)
	h.UpdateMovingAvg(chosenArm, reward)
	h.updateTimeSinceLast(chosenArm)
}

func (h *History) updateRounds(chosenArm *bandit.Arm, armToProbability map[*bandit.Arm]float64, reward float64) {
	for arm, probability := range armToProbability {
		armStats := h.ArmToStats[arm]
		armStats.Rounds = append(armStats.Rounds, &Round{
			Probability: probability,
			Reward:      reward,
			WasChosen:   arm == chosenArm,
		})
	}
}

func (h *History) UpdateMovingAvg(chosenArm *bandit.Arm, reward float64) {
	armStats := h.ArmToStats[chosenArm]
	armStats.Count++
	count := float64(armStats.Count)
	armStats.MovingAvg = ((count-1.0)/count)*armStats.MovingAvg + (1.0/count)*reward
}

func (h *History) updateAvgRewards(armToProbability map[*bandit.Arm]float64, reward float64) {
	for arm := range armToProbability {
		armStats := h.ArmToStats[arm]
		avgRewardWhenUsed, avgRewardWhenEligible := h.calculateAvgRewards(armStats.Rounds)
		armStats.AvgRewardWhenUsed = avgRewardWhenUsed
		armStats.AvgRewardWhenEligible = avgRewardWhenEligible
	}
}

func (h *History) calculateAvgRewards(rounds []*Round) (float64, float64) {
	weightedProbSumWhenUsed := 0.0
	probSumWhenUsed := 0.0
	weightedProbSumWhenEligible := 0.0
	probSumWhenEligible := 0.0
	for _, round := range rounds {
		if round.WasChosen {
			weightedProbSumWhenUsed += round.Probability * round.Reward
			probSumWhenUsed += round.Probability
		}
		weightedProbSumWhenEligible += round.Probability * round.Reward
		probSumWhenEligible += round.Probability

	}
	return weightedProbSumWhenUsed / probSumWhenUsed, weightedProbSumWhenEligible / probSumWhenEligible
}

func (h *History) updateTimeSinceLast(chosenArm *bandit.Arm) {
	for arm, stats := range h.ArmToStats {
		if arm == chosenArm {
			stats.TimeSinceLast = 0
		} else {
			stats.TimeSinceLast++
		}
	}
}
