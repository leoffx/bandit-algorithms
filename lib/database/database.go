package database

import (
	"github.com/leoffx/bandit-algorithms/lib/bandit"
)

type Entry struct {
	Round        int
	ChosenArm    *bandit.Arm
	EligibleArms []*bandit.Arm
	Reward       float64
}

type Database struct {
	Entries []*Entry
}

func NewDatabase() *Database {
	return &Database{}
}

func NewEntry(round int, chosenArm *bandit.Arm, eligibleArms []*bandit.Arm, reward float64) *Entry {
	return &Entry{
		Round:        round,
		ChosenArm:    chosenArm,
		EligibleArms: eligibleArms,
		Reward:       reward,
	}
}

func (d *Database) Insert(e *Entry) {
	d.Entries = append(d.Entries, e)
}
