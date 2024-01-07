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
	armToStats := &database.ArmToStats{
		arms[0]: database.NewArmStats(1, 5),
		arms[1]: database.NewArmStats(1, 0),
		arms[2]: database.NewArmStats(1, 10),
	}

	tests := []struct {
		name       string
		strategy   *strategy.EpsilonGreedy
		arms       []*bandit.Arm
		armToStats *database.ArmToStats
		expected   *bandit.Arm
	}{
		{
			name:       "Choose best arm with clear winner",
			strategy:   strategy.NewEpsilonGreedy(0),
			arms:       arms,
			armToStats: armToStats,
			expected:   arms[2],
		},
		{
			name:     "Choose best arm with tie",
			strategy: strategy.NewEpsilonGreedy(0),
			arms: []*bandit.Arm{
				bandit.NewArm(10, 0),
				bandit.NewArm(10, 0),
			},
			armToStats: &database.ArmToStats{
				arms[0]: database.NewArmStats(1, 10),
				arms[1]: database.NewArmStats(1, 10),
			},
			expected: arms[0],
		},
		{
			name:       "Choose best arm with no entries",
			strategy:   strategy.NewEpsilonGreedy(0),
			arms:       arms,
			armToStats: &database.ArmToStats{},
			expected:   arms[0],
		},
		{
			name:       "Choose random arm with epsilon 0.5",
			strategy:   strategy.NewEpsilonGreedy(0.5),
			arms:       arms,
			armToStats: armToStats,
			expected:   arms[1],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arm := tt.strategy.ChooseArm(tt.arms, tt.armToStats)
			if arm != tt.expected {
				t.Errorf("Test '%s' failed: Expected %v, got %v", tt.name, tt.expected, arm)
			}
		})
	}
}
