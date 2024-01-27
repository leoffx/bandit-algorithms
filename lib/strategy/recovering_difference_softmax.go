package strategy

import (
	"errors"

	"github.com/leoffx/bandit-algorithms/lib/arm"
	"github.com/leoffx/bandit-algorithms/lib/database"
)

type RecoveringDifferenceSoftmax struct {
	temperature float64
}

func NewRecoveringDifferenceSoftmax(temperature float64) *RecoveringDifferenceSoftmax {
	if temperature <= 0 {
		panic(errors.New("temperature must be positive"))
	}
	return &RecoveringDifferenceSoftmax{
		temperature: temperature,
	}
}

func (r *RecoveringDifferenceSoftmax) ChooseArm(armToScore *database.ArmToScore) arm.Arm {
	ks := make([]arm.Arm, len(*armToScore))
	vs := make([]float64, len(*armToScore))
	i := 0
	for k, v := range *armToScore {
		ks[i] = k
		vs[i] = v
		i++
	}
	arm, err := RandomChoice(ks, vs)
	if err != nil {
		panic(err)
	}
	return *arm
}

func (r *RecoveringDifferenceSoftmax) ScoreArms(arms []arm.Arm, armToStats *database.ArmToStats) *database.ArmToScore {
	logits := make([]float64, len(arms))
	for i, arm := range arms {
		armStats := (*armToStats)[arm]
		if armStats == nil {
			logits[i] = 1 / float64(len(arms))
			continue
		}
		logits[i] = (armStats.AvgRewardWhenUsed - armStats.AvgRewardWhenEligible) / armStats.AvgRewardWhenEligible
	}

	probabilities := Softmax(logits, r.temperature)
	armToScore := make(database.ArmToScore)
	for i, arm := range arms {
		armToScore[arm] = probabilities[i]
	}
	return &armToScore
}
