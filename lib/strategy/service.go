package strategy

import (
	"github.com/leoffx/bandit-algorithms/lib/bandit"
	"github.com/leoffx/bandit-algorithms/lib/database"
)

const (
	EpsilonGreedyStrategy               = iota
	RecoveringDifferenceSoftmaxStrategy = iota
)

type Strategy interface {
	ChooseArm([]*bandit.Arm, []*database.Entry) *bandit.Arm
}
