package bandit

import (
	"math/rand"
	"strconv"
)

type Bandit struct {
	Arms []*Arm
}

type Arm struct {
	stdDev float64
	mean   float64
}

func NewBandit(numArms int) *Bandit {
	arms := InitializeArms(numArms)
	return &Bandit{
		Arms: arms,
	}
}

func NewArm(mean float64, stdDev float64) *Arm {
	return &Arm{
		mean:   mean,
		stdDev: stdDev,
	}
}

func (b Bandit) PullArm(a *Arm) float64 {
	return rand.NormFloat64()*a.stdDev + a.mean
}

func (a Arm) String() string {
	return "Arm{mean: " + strconv.FormatFloat(a.mean, 'f', -1, 64) + ", stdDev: " + strconv.FormatFloat(a.stdDev, 'f', -1, 64) + "}"
}

func (b Bandit) GetEligibleArms() []*Arm {
	// TODO: Implement
	return b.Arms
}

func InitializeArms(numArms int) []*Arm {
	arms := make([]*Arm, numArms)
	for i := 0; i < numArms; i++ {
		arms[i] = NewArm(rand.Float64(), rand.Float64())
	}
	return arms
}
