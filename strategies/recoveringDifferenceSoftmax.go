package strategies

import (
	"errors"
	"math"
	"math/rand"

	"github.com/leoffx/a-sleeping-recovering-bandit-algorithm-for-optimizing-recurring-notifications/structs"
)

type RecoveringDifferenceSoftmax struct {
	history     structs.History
	temperature float64
}

func NewRecoveringDifferenceSoftmax(history structs.History, temperature float64) (*RecoveringDifferenceSoftmax, error) {
	if temperature <= 0 {
		return nil, errors.New("temperature must be positive")
	}
	return &RecoveringDifferenceSoftmax{
		history:     history,
		temperature: temperature,
	}, nil
}

func (r *RecoveringDifferenceSoftmax) CalculateArmsProbabilities(arms []*structs.Arm) map[*structs.Arm]float64 {
	scores := r.calculateScores(arms)
	probabilities := softmax(scores, r.temperature)
	armToProbability := make(map[*structs.Arm]float64)
	for i, arm := range arms {
		armToProbability[arm] = probabilities[i]
	}
	return armToProbability
}

func (r *RecoveringDifferenceSoftmax) ChooseArm(armToProbability map[*structs.Arm]float64) *structs.Arm {
	arms := make([]*structs.Arm, 0, len(armToProbability))
	probabilities := make([]float64, 0, len(armToProbability))
	for arm, probability := range armToProbability {
		arms = append(arms, arm)
		probabilities = append(probabilities, probability)
	}
	return randomChoices(arms, probabilities)
}

func (r *RecoveringDifferenceSoftmax) calculateScores(arms []*structs.Arm) []float64 {
	scores := make([]float64, len(arms))
	for i, arm := range arms {
		armStats := r.history.ArmToStats[arm]
		scores[i] = (armStats.AvgRewardWhenUsed - armStats.AvgRewardWhenEligible) / armStats.AvgRewardWhenEligible
	}
	return scores
}

func softmax(logits []float64, temperature float64) []float64 {
	expSum := 0.0
	for _, logit := range logits {
		expSum += math.Exp(logit / temperature)
	}

	probabilities := make([]float64, len(logits))
	for i, logit := range logits {
		probabilities[i] = math.Exp(logit/temperature) / expSum
	}

	return probabilities
}

func randomChoices(arms []*structs.Arm, probabilities []float64) *structs.Arm {
	r := rand.Float64()
	for i, probability := range probabilities {
		r -= probability
		if r <= 0 {
			return arms[i]
		}
	}
	return arms[len(arms)-1]
}
