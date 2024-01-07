package database

import (
	"math"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
)

type DatabaseAggregator struct {
	*Database
}

type ArmStats struct {
	Count     int
	AvgReward float64
}

type ArmToStats map[*bandit.Arm]*ArmStats

func NewDatabaseAggregator() *DatabaseAggregator {
	db := NewDatabase()
	return &DatabaseAggregator{
		Database: db,
	}
}

func (db *DatabaseAggregator) ArmToStats() ArmToStats {
	armToStats := make(ArmToStats)
	negInf := math.Inf(-1)
	for _, entry := range db.Entries {
		curr, found := armToStats[entry.ChosenArm]
		if !found {
			curr = &ArmStats{
				Count:     0,
				AvgReward: negInf,
			}
			armToStats[entry.ChosenArm] = curr
		}
		curr.Count++
		if curr.AvgReward == negInf {
			curr.AvgReward = entry.Reward
		} else {
			newAvg := curr.AvgReward + (entry.Reward-curr.AvgReward)/float64(curr.Count)
			curr.AvgReward = newAvg
		}
	}
	return armToStats
}
