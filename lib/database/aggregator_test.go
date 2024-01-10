package database_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
	"github.com/leoffx/bandit-algorithms/lib/database"
)

func TestArmToStats(t *testing.T) {
	bandit := bandit.NewBandit(3)
	arm0 := bandit.Arms[0]
	arm1 := bandit.Arms[1]
	arm2 := bandit.Arms[2]

	db := database.NewDatabaseAggregator()
	db.AddEntry(database.NewEntry(0, arm0, &database.ArmToScore{arm0: 0.9, arm1: 0.1}, 0.5))
	db.AddEntry(database.NewEntry(1, arm1, &database.ArmToScore{arm0: 0.5, arm1: 0.5}, 1))
	db.AddEntry(database.NewEntry(2, arm1, &database.ArmToScore{arm0: 0.25, arm1: 0.25}, 1))
	db.AddEntry(database.NewEntry(3, arm0, &database.ArmToScore{arm0: 0.5, arm1: 0.5}, 0))
	db.AddEntry(database.NewEntry(4, arm1, &database.ArmToScore{arm0: 0.25, arm1: 0.5, arm2: 0.25}, 1))

	armToStats := db.ArmToStats()
	arm0Stats := armToStats[arm0]
	arm1Stats := armToStats[arm1]
	arm2Stats := armToStats[arm2]
	// AvgRewardWhenUsed
	if arm0Stats.AvgRewardWhenUsed != 0.25 {
		t.Error("arm0Stats.AvgRewardWhenUsed != 0.25")
	}
	if arm1Stats.AvgRewardWhenUsed != 1 {
		t.Error("arm1Stats.AvgRewardWhenUsed != 1")
	}
	if arm2Stats.AvgRewardWhenUsed != 0 {
		t.Error("arm2Stats.AvgRewardWhenUsed != 0")
	}
	// AvgRewardWhenElligible
	if arm0Stats.AvgRewardWhenElligible != 29./48. {
		t.Error("arm0Stats.AvgRewardWhenElligible != 29./48.")
	}
	if math.Abs(arm1Stats.AvgRewardWhenElligible-26./37.) > 1e-6 {
		t.Error("arm1Stats.AvgRewardWhenElligible != 26./37.")
	}
	if arm2Stats.AvgRewardWhenElligible != 1 {
		t.Error("arm2Stats.AvgRewardWhenElligible != 1")
	}
	fmt.Println(armToStats)
}
