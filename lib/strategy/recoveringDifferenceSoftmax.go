package strategy

import (
	"errors"
	"math"
	"math/rand"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
	"github.com/leoffx/bandit-algorithms/lib/database"
)

type RecoveringDifferenceSoftmax struct {
	history     *database.History
	temperature float64
}

func NewRecoveringDifferenceSoftmax(history *database.History, temperature float64) (*RecoveringDifferenceSoftmax, error) {
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
	probabilities := softmax(scores, r.temperature)
	armToProbability := make(map[*bandit.Arm]float64)
	for i, arm := range arms {
		armToProbability[arm] = probabilities[i]
	}
	return armToProbability
}

func (r *RecoveringDifferenceSoftmax) ChooseArm(armToProbability map[*bandit.Arm]float64) *bandit.Arm {
	arms := make([]*bandit.Arm, 0, len(armToProbability))
	probabilities := make([]float64, 0, len(armToProbability))
	for arm, probability := range armToProbability {
		arms = append(arms, arm)
		probabilities = append(probabilities, probability)
	}
	return randomChoices(arms, probabilities)
}

func (r *RecoveringDifferenceSoftmax) calculateScores(arms []*bandit.Arm) []float64 {
	scores := make([]float64, len(arms))
	for i, arm := range arms {
		armStats := r.history.ArmToStats[arm]
		if armStats.Count == 0 {
			scores[i] = 0
			continue
		}
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

func randomChoices(arms []*bandit.Arm, probabilities []float64) *bandit.Arm {
	r := rand.Float64()
	for i, probability := range probabilities {
		r -= probability
		if r <= 0 {
			return arms[i]
		}
	}
	return arms[len(arms)-1]
}
