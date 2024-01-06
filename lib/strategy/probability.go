package strategy

import (
	"errors"
	"math"
	"math/rand"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
)

func Softmax(logits []float64, temperature float64) []float64 {
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

func RandomChoices(arms []*bandit.Arm, probabilities *[]float64) (*bandit.Arm, error) {
	if len(arms) == 0 {
		return nil, errors.New("no arms to choose from")
	}
	if probabilities == nil || len(*probabilities) == 0 {
		return arms[rand.Intn(len(arms))], nil
	}
	if len(*probabilities) != len(arms) {
		return nil, errors.New("probabilities must be the same length as arms")
	}
	r := rand.Float64()
	for i, probability := range *probabilities {
		r -= probability
		if r <= 0 {
			return arms[i], nil
		}
	}
	return arms[len(arms)-1], nil
}
