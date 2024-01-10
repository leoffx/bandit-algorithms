package bandit

import (
	"github.com/leoffx/bandit-algorithms/lib/arm"
)

type Bandit struct {
	Arms []arm.Arm
}

func NewBandit(numArms int) *Bandit {
	arms := arm.InitializeArms(numArms)
	return &Bandit{
		Arms: arms,
	}
}

func (b Bandit) PullArm(a arm.Arm) float64 {
	return a.Pull()
}

func (b Bandit) String() string {
	arms := ""
	for _, a := range b.Arms {
		arms += a.String() + "\n"
	}
	return arms
}

func (b Bandit) GetEligibleArms() []arm.Arm {
	// TODO: Implement
	return b.Arms
}
