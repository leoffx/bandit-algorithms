package strategies

import (
	"errors"
	"math/rand"

	"github.com/leoffx/a-sleeping-recovering-bandit-algorithm-for-optimizing-recurring-notifications/structs"
)

type EpsilonGreedy struct {
	history structs.History
	epsilon float64
}

func NewEpsilonGreedy(history structs.History, epsilon float64) (Strategy, error) {
	if epsilon < 0.0 || epsilon > 1.0 {
		return &EpsilonGreedy{}, errors.New("epsilon must be between 0.0 and 1.0")
	}
	return &EpsilonGreedy{history, epsilon}, nil
}

func (e *EpsilonGreedy) ChooseArm(arms []structs.Arm) structs.Arm {
	if rand.Float64() > e.epsilon {
		return e.chooseBestArm(arms)
	}
	return e.chooseRandomArm(arms)
}

func (e *EpsilonGreedy) chooseBestArm(arms []structs.Arm) structs.Arm {
	var bestArm structs.Arm
	var bestArmValue float64 = 0.0

	for _, arm := range arms {
		armStats := e.history.ArmToStats[arm]
		if armStats.Value < bestArmValue {
			continue
		}
		bestArm = arm
		bestArmValue = armStats.Value
	}
	return bestArm
}

func (e *EpsilonGreedy) chooseRandomArm(arms []structs.Arm) structs.Arm {
	return arms[rand.Intn(len(arms))]
}
