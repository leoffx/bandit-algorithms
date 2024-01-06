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

func (e *EpsilonGreedy) ChooseArm(arms []*bandit.Arm, entries []*database.Entry) *bandit.Arm {
	// Explore
	if rand.Float64() < e.epsilon || len(entries) == 0 {
		arm, err := RandomChoices(arms, nil)
		if err != nil {
			panic(err)
		}
		return arm
	}
	// Exploit
	return e.chooseBestArm(entries)
}

type ArmStats struct {
	Count     int
	AvgReward *float64
}

func (e *EpsilonGreedy) chooseBestArm(entries []*database.Entry) *bandit.Arm {
	armToReward := make(map[*bandit.Arm]*ArmStats)
	var bestArm *bandit.Arm
	var bestAvgReward *float64
	for _, entry := range entries {
		curr, found := armToReward[entry.ChosenArm]
		if !found {
			curr = &ArmStats{
				Count:     0,
				AvgReward: nil,
			}
			armToReward[entry.ChosenArm] = curr
		}
		curr.Count++
		if curr.AvgReward == nil {
			avg := entry.Reward
			curr.AvgReward = &avg
		} else {
			newAvg := *curr.AvgReward + (entry.Reward-*curr.AvgReward)/float64(curr.Count)
			curr.AvgReward = &newAvg
		}
		if bestAvgReward == nil || *curr.AvgReward > *bestAvgReward {
			bestArm = entry.ChosenArm
			bestAvgReward = curr.AvgReward
		}
	}
	return bestArm
}
