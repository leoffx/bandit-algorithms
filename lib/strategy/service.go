package strategy

import "github.com/leoffx/bandit-algorithms/lib/bandit"

const (
	EpsilonGreedyStrategy               = iota
	RecoveringDifferenceSoftmaxStrategy = iota
)

type Strategy interface {
	ChooseArm(map[*bandit.Arm]float64) *bandit.Arm
	CalculateArmsProbabilities([]*bandit.Arm) map[*bandit.Arm]float64
}
