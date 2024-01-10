package arm

import (
	"fmt"
	"math/rand"
)

type BernouliArm struct {
	probability float64
}

func NewBernouliArm() *BernouliArm {
	probability := rand.Float64()
	return &BernouliArm{
		probability: probability,
	}
}

func (a BernouliArm) Pull() float64 {
	if rand.Float64() < a.probability {
		return 1
	}
	return 0
}

func (a BernouliArm) String() string {
	return fmt.Sprintf("BernouliArm{probability: %f}", a.probability)

}

func (a BernouliArm) Mean() float64 {
	return a.probability
}
