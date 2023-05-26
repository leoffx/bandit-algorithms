package strategies

import (
	"github.com/leoffx/a-sleeping-recovering-bandit-algorithm-for-optimizing-recurring-notifications/structs"
)

const (
	EpsilonGreedyStrategy               = iota
	RecoveringDifferenceSoftmaxStrategy = iota
)

type Strategy interface {
	ChooseArm([]*structs.Arm) *structs.Arm
}
