package database

import (
	"github.com/leoffx/bandit-algorithms/lib/bandit"
)

type DatabaseAggregator struct {
	*Database
}

type ArmStats struct {
	count             int
	AvgRewardWhenUsed float64
	// avgRewardWhenElligible float64
}

type ArmToStats map[*bandit.Arm]*ArmStats

func NewArmStats(count int, avgRewardWhenUsed float64) *ArmStats {
	return &ArmStats{
		count:             count,
		AvgRewardWhenUsed: avgRewardWhenUsed,
	}
}

func NewDatabaseAggregator() *DatabaseAggregator {
	db := NewDatabase()
	return &DatabaseAggregator{
		Database: db,
	}
}

func (db *DatabaseAggregator) ArmToStats() ArmToStats {
	armToStats := make(ArmToStats)
	for _, entry := range db.Entries {
		curr := armToStats[entry.ChosenArm]

		if curr == nil {
			curr = &ArmStats{}
			armToStats[entry.ChosenArm] = curr
		}

		curr.count++
		curr.AvgRewardWhenUsed = curr.AvgRewardWhenUsed + (entry.Reward-curr.AvgRewardWhenUsed)/float64(curr.count)

		// for _, arm := range entry.EligibleArms {
		// 	curr := armToStats[arm]

		// 	if curr == nil {
		// 		curr = &ArmStats{}
		// 		armToStats[arm] = curr
		// 	}

		// 	curr.avgRewardWhenElligible = curr.avgRewardWhenElligible + (entry.Reward-curr.avgRewardWhenElligible)/float64(curr.count)
		// }
	}
	return armToStats
}
