package strategy

import (
	"github.com/leoffx/bandit-algorithms/lib/arm"
	"github.com/leoffx/bandit-algorithms/lib/database"
)

const (
	EpsilonGreedyStrategy               = iota
	RecoveringDifferenceSoftmaxStrategy = iota
)

type Strategy interface {
	ChooseArm(*database.ArmToScore) arm.Arm
	ScoreArms([]arm.Arm, *database.ArmToStats) *database.ArmToScore
}
