package database

import (
	"github.com/leoffx/bandit-algorithms/lib/arm"
)

type DatabaseAggregator struct {
	*Database
	lastComputedRound int
	armToStats        ArmToStats
}

type ArmStats struct {
	count                  int
	AvgRewardWhenUsed      float64
	AvgRewardWhenElligible float64
	sumScoresWhenElligible float64
}

type ArmToStats map[arm.Arm]*ArmStats

func NewArmStats(count int, avgRewardWhenUsed float64) *ArmStats {
	return &ArmStats{
		count:             count,
		AvgRewardWhenUsed: avgRewardWhenUsed,
	}
}

func NewDatabaseAggregator() *DatabaseAggregator {
	db := NewDatabase()
	return &DatabaseAggregator{
		Database:   db,
		armToStats: make(ArmToStats),
	}
}

func (db *DatabaseAggregator) ArmToStats() *ArmToStats {
	for _, entry := range db.Entries {
		if entry.Round >= db.lastComputedRound {
			updateChosenArm(db.armToStats, entry)
			updateAllArms(db.armToStats, entry)
			db.lastComputedRound = entry.Round
		}
	}
	return &db.armToStats
}

func updateChosenArm(armToStats ArmToStats, entry *Entry) {
	curr := armToStats[entry.ChosenArm]
	if curr == nil {
		curr = &ArmStats{}
		armToStats[entry.ChosenArm] = curr
	}
	curr.count++
	curr.AvgRewardWhenUsed = curr.AvgRewardWhenUsed + (entry.Reward-curr.AvgRewardWhenUsed)/float64(curr.count)
}

func updateAllArms(armToStats ArmToStats, entry *Entry) {
	for arm, score := range *entry.ArmToScore {
		curr := armToStats[arm]
		if curr == nil {
			curr = &ArmStats{}
			armToStats[arm] = curr
		}

		// sum (score * reward) / sum (score)
		dividend := curr.sumScoresWhenElligible + score
		curr.AvgRewardWhenElligible = (curr.AvgRewardWhenElligible * curr.sumScoresWhenElligible / dividend) + (score * entry.Reward / dividend)
		curr.sumScoresWhenElligible += score
	}
}
