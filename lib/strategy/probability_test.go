package strategy_test

import (
	"math"
	"math/rand"
	"reflect"
	"testing"

	"github.com/leoffx/bandit-algorithms/lib/arm"
	"github.com/leoffx/bandit-algorithms/lib/strategy"
)

func TestSoftmax(t *testing.T) {
	tests := []struct {
		logits      []float64
		temperature float64
		expected    []float64
	}{
		{
			logits:      []float64{1.0, 2.0, 3.0},
			temperature: 1,
			expected:    []float64{0.09003057, 0.24472847, 0.66524096},
		},
		{
			logits:      []float64{1.0, 1.0, 1.0},
			temperature: 3.0,
			expected:    []float64{0.33333333, 0.33333333, 0.33333333},
		},
		{
			logits:      []float64{2.0, 1.0, 0.1},
			temperature: 0.5,
			expected:    []float64{0.86377712, 0.11689952, 0.01932336},
		},
	}
	for i, tt := range tests {
		got := strategy.Softmax(tt.logits, tt.temperature)
		for j, v := range got {
			if math.Abs(v-tt.expected[j]) > 1e-6 {
				t.Errorf("TestSoftmax(%d): got %v, want %v", i, got, tt.expected)
				break
			}
		}
	}

}

func TestRandomChoices(t *testing.T) {
	rand.Seed(42)
	tests := []struct {
		arms          []arm.Arm
		probabilities []float64
		expected      arm.Arm
	}{
		{
			arms: []arm.Arm{
				arm.NewGaussianArm(0.1, 0.1),
				arm.NewGaussianArm(0.2, 0.1),
				arm.NewGaussianArm(0.3, 0.1),
			},
			probabilities: []float64{0, 1, 0},
			expected:      arm.NewGaussianArm(0.2, 0.1),
		},
		{
			arms: []arm.Arm{
				arm.NewGaussianArm(0.1, 0.1),
				arm.NewGaussianArm(0.2, 0.1),
				arm.NewGaussianArm(0.3, 0.1),
			},
			probabilities: nil,
			expected:      arm.NewGaussianArm(0.3, 0.1),
		},
	}
	for i, tt := range tests {
		got, _ := strategy.RandomChoice(tt.arms, tt.probabilities)
		if !reflect.DeepEqual(*got, tt.expected) {
			t.Errorf("TestRandomChoices(%d): got %v, want %v", i, got, tt.expected)
		}
	}
}
