package structs

import (
	"math/rand"
	"strconv"
)

type Arm struct {
	stdDev float64
	mean   float64
}

// Draw from a normal distribution with mean and stdDev
func (a Arm) DrawReward() float64 {
	return rand.NormFloat64()*a.stdDev + a.mean
}

func (a Arm) String() string {
	return "Arm{mean: " + strconv.FormatFloat(a.mean, 'f', -1, 64) + ", stdDev: " + strconv.FormatFloat(a.stdDev, 'f', -1, 64) + "}"
}

func InitializeArms(numArms int) []*Arm {
	var arms []*Arm
	for i := 0; i < numArms; i++ {
		arm := &Arm{
			mean:   rand.Float64(),
			stdDev: rand.Float64(),
		}
		arms = append(arms, arm)
	}
	return arms
}

func GetEligibleArms(arms []*Arm) []*Arm {
	// TODO: Implement
	return arms
}
