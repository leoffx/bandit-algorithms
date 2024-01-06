package strategy_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
	"github.com/leoffx/bandit-algorithms/lib/strategy"
)

func TestSoftmax(t *testing.T) {
	// ...
}

func TestRandomChoices(t *testing.T) {
	rand.Seed(42)
	tests := []struct {
		arms          []*bandit.Arm
		probabilities []float64
		expected      *bandit.Arm
	}{
		{
			arms: []*bandit.Arm{
				bandit.NewArm(0.1, 0.1),
				bandit.NewArm(0.2, 0.1),
				bandit.NewArm(0.3, 0.1),
			},
			probabilities: []float64{0, 1, 0},
			expected:      bandit.NewArm(0.2, 0.1),
		},
		{
			arms: []*bandit.Arm{
				bandit.NewArm(0.1, 0.1),
				bandit.NewArm(0.2, 0.1),
				bandit.NewArm(0.3, 0.1),
			},
			probabilities: nil,
			expected:      bandit.NewArm(0.3, 0.1),
		},
	}
	for i, tt := range tests {
		got, _ := strategy.RandomChoices(tt.arms, &tt.probabilities)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("TestRandomChoices(%d): got %v, want %v", i, got, tt.expected)
		}
	}
}
