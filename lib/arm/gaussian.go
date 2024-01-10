package arm

import (
	"fmt"
	"math/rand"
)

type GaussianArm struct {
	mean   float64
	stdDev float64
}

func NewGaussianArm(mean float64, stdDev float64) *GaussianArm {
	return &GaussianArm{
		mean:   mean,
		stdDev: stdDev,
	}
}

func (a GaussianArm) Pull() float64 {
	return rand.NormFloat64()*a.stdDev + a.mean
}

func (a GaussianArm) String() string {
	return fmt.Sprintf("GaussianArm{mean: %f, stdDev: %f}", a.mean, a.stdDev)
}

func (a GaussianArm) Mean() float64 {
	return a.mean
}
