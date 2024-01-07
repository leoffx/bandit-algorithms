package strategy

import (
	"errors"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
	"github.com/leoffx/bandit-algorithms/lib/database"
)

type RecoveringDifferenceSoftmax struct {
	history     *database.Database
	temperature float64
}

func NewRecoveringDifferenceSoftmax(history *database.Database, temperature float64) (*RecoveringDifferenceSoftmax, error) {
	if temperature <= 0 {
		return nil, errors.New("temperature must be positive")
	}
	return &RecoveringDifferenceSoftmax{
		history:     history,
		temperature: temperature,
	}, nil
}

func (r *RecoveringDifferenceSoftmax) CalculateArmsProbabilities(arms []*bandit.Arm) map[*bandit.Arm]float64 {
	scores := r.calculateScores(arms)
	probabilities := Softmax(scores, r.temperature)
	armToProbability := make(map[*bandit.Arm]float64)
	for i, arm := range arms {
		armToProbability[arm] = probabilities[i]
	}
	return armToProbability
}

func (r *RecoveringDifferenceSoftmax) ChooseArm(arms []*bandit.Arm, entries []*database.Entry) *bandit.Arm {
	// arms := make([]*bandit.Arm, 0, len(armToProbability))
	// probabilities := make([]float64, 0, len(armToProbability))
	// for arm, probability := range armToProbability {
	// 	arms = append(arms, arm)
	// 	probabilities = append(probabilities, probability)
	// }
	// return randomChoices(arms, probabilities)
	return nil
}

func (r *RecoveringDifferenceSoftmax) calculateScores(arms []*bandit.Arm) []float64 {
	// scores := make([]float64, len(arms))
	// for i, arm := range arms {
	// 	armStats := r.history.ArmToStats[arm]
	// 	if armStats.Count == 0 {
	// 		scores[i] = 0
	// 		continue
	// 	}
	// 	scores[i] = (armStats.AvgRewardWhenUsed - armStats.AvgRewardWhenEligible) / armStats.AvgRewardWhenEligible
	// }
	// return scores
	return nil
}