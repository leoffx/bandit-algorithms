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
		arm := &Arm{
			mean:   rand.Float64(),
			stdDev: rand.Float64(),
		}
		arms[i] = arm
	}
	return arms
}
