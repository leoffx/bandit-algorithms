package strategy_test

import (
	"math/rand"
	"testing"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
	"github.com/leoffx/bandit-algorithms/lib/database"
	"github.com/leoffx/bandit-algorithms/lib/strategy"
)

func TestChooseArm(t *testing.T) {
	rand.Seed(42)
	arms := []*bandit.Arm{
		bandit.NewArm(5, 0),
		bandit.NewArm(0, 0),
		bandit.NewArm(10, 0),
	}
	entries := []*database.Entry{
		database.NewEntry(1, arms[0], arms, 5),
		database.NewEntry(2, arms[1], arms, 0),
		database.NewEntry(3, arms[2], arms, 10),
	}

	tests := []struct {
		name     string
		strategy *strategy.EpsilonGreedy
		arms     []*bandit.Arm
		entries  []*database.Entry
		expected *bandit.Arm
	}{
		{
			name:     "Choose best arm with clear winner",
			strategy: strategy.NewEpsilonGreedy(0),
			arms:     arms,
			entries:  entries,
			expected: arms[2],
		},
		{
			name:     "Choose best arm with tie",
			strategy: strategy.NewEpsilonGreedy(0),
			arms: []*bandit.Arm{
				bandit.NewArm(10, 0),
				bandit.NewArm(10, 0),
			},
			entries: []*database.Entry{
				database.NewEntry(1, arms[0], arms, 10),
				database.NewEntry(2, arms[1], arms, 10),
			},
			expected: arms[0],
		},
		{
			name:     "Choose best arm with no entries",
			strategy: strategy.NewEpsilonGreedy(0),
			arms:     arms,
			entries:  []*database.Entry{},
			expected: arms[0],
		},
		{
			name:     "Choose random arm with epsilon 0.5",
			strategy: strategy.NewEpsilonGreedy(0.5),
			arms:     arms,
			entries:  entries,
			expected: arms[1],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arm := tt.strategy.ChooseArm(tt.arms, tt.entries)
			if arm != tt.expected {
				t.Errorf("Test '%s' failed: Expected %v, got %v", tt.name, tt.expected, arm)
			}
		})
	}
}
