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

func (e *EpsilonGreedy) ChooseArm(arms []*bandit.Arm, armToStats *database.ArmToStats) *bandit.Arm {
	// Explore
	if rand.Float64() < e.epsilon || len(*armToStats) == 0 {
		arm, err := RandomChoices(arms, nil)
		if err != nil {
			panic(err)
		}
		return arm
	}
	// Exploit
	return e.chooseBestArm(armToStats)
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
