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

	probs := make([]float64, len(logits))
	for i, logit := range logits {
		probs[i] = math.Exp(logit/temperature) / expSum
	}

	return probs
}

func RandomChoice[T any](items []T, probs []float64) (*T, error) {
	if len(items) == 0 {
		return nil, errors.New("no items to choose from")
	}
	if len(probs) == 0 {
		return &items[rand.Intn(len(items))], nil
	}
	if len(probs) != len(items) {
		return nil, errors.New("probs must be the same length as items")
	}
	probsSum := 0.0
	for _, prob := range probs {
		probsSum += prob
	}
	if math.Abs(probsSum-1) > 1e-6 {
		return nil, errors.New("probs must sum to 1")
	}

	r := rand.Float64()
	for i, prob := range probs {
		r -= prob
		if r <= 0 {
			return &items[i], nil
		}
	}
	return &items[len(items)-1], nil
}
