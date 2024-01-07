package strategy

import (
	"math/rand"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
	"github.com/leoffx/bandit-algorithms/lib/database"
)

type EpsilonGreedy struct {
	epsilon float64
}

func NewEpsilonGreedy(epsilon float64) *EpsilonGreedy {
	return &EpsilonGreedy{
		epsilon: epsilon,
	}
}

func (e *EpsilonGreedy) ScoreArms(arms []*bandit.Arm, armToStats *database.ArmToStats) *database.ArmToScore {
	armToScore := make(database.ArmToScore)
	// Explore
	if rand.Float64() < e.epsilon || len(*armToStats) == 0 {
		for _, arm := range arms {
			armToScore[arm] = 1 / float64(len(arms))
		}
		return &armToScore
	}
	// Exploit
	bestArm := e.chooseBestArm(armToStats)
	for _, arm := range arms {
		if arm == bestArm {
			armToScore[arm] = 1
		} else {
			armToScore[arm] = 0
		}
	}
	return &armToScore
}

func (e *EpsilonGreedy) ChooseArm(armToScore *database.ArmToScore) *bandit.Arm {
	ks := make([]*bandit.Arm, 0, len(*armToScore))
	vs := make([]float64, 0, len(*armToScore))
	for k, v := range *armToScore {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	arm, err := RandomChoice(ks, &vs)
	if err != nil {
		panic(err)
	}
	return arm
}

func (e *EpsilonGreedy) chooseBestArm(armToStats *database.ArmToStats) *bandit.Arm {
	var bestArm *bandit.Arm
	var bestAvgReward float64
	for arm, stats := range *armToStats {
		if bestArm == nil || stats.AvgRewardWhenUsed > bestAvgReward {
			bestArm = arm
			bestAvgReward = stats.AvgRewardWhenUsed
		}
	}
	return bestArm
}
