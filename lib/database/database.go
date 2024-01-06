package database

import (
	"fmt"

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

func (d *Database) String() string {
	entries := ""
	for _, e := range d.Entries {
		entries += e.String() + "\n"
	}
	return entries
}

func (e *Entry) String() string {
	return fmt.Sprintf("Round: %d ChosenArm: %s Reward: %g", e.Round, e.ChosenArm, e.Reward)
}
