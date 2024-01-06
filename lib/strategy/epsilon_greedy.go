package strategy

import (
	"math/rand"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
	"github.com/leoffx/bandit-algorithms/lib/database"
)

type EpsilonGreedy struct {
	epsilon float64
}

func NewEpsilonGreedy() *EpsilonGreedy {
	epsilon := 0.1
	return &EpsilonGreedy{
		epsilon: epsilon,
	}
}

func (e *EpsilonGreedy) ChooseArm(arms []*bandit.Arm, entries []*database.Entry) *bandit.Arm {
	if rand.Float64() < e.epsilon {
		// Explore
		RandomChoices(arms, nil)
	}
	// Exploit
	return arms[0]
}
