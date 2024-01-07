package strategy

import (
	"errors"
	"math"
	"math/rand"
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

func RandomChoice[T any](items []*T, probabilities *[]float64) (*T, error) {
	if len(items) == 0 {
		return nil, errors.New("no items to choose from")
	}
	if probabilities == nil || len(*probabilities) == 0 {
		return items[rand.Intn(len(items))], nil
	}
	if len(*probabilities) != len(items) {
		return nil, errors.New("probabilities must be the same length as items")
	}
	probabilitiesSum := 0.0
	for _, probability := range *probabilities {
		probabilitiesSum += probability
	}
	if math.Abs(probabilitiesSum-1) > 1e-6 {
		return nil, errors.New("probabilities must sum to 1")
	}

	r := rand.Float64()
	for i, probability := range *probabilities {
		r -= probability
		if r <= 0 {
			return items[i], nil
		}
	}
	return items[len(items)-1], nil
}
