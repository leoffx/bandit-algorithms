package database

import (
	"fmt"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
)

type ArmToScore map[*bandit.Arm]float64

type Entry struct {
	Round      int
	ChosenArm  *bandit.Arm
	ArmToScore *ArmToScore
	Reward     float64
}

type Database struct {
	Entries []*Entry
}

func NewDatabase() *Database {
	return &Database{}
}

func NewEntry(round int, chosenArm *bandit.Arm, armToScore *ArmToScore, reward float64) *Entry {
	return &Entry{
		Round:      round,
		ChosenArm:  chosenArm,
		ArmToScore: armToScore,
		Reward:     reward,
	}
}

func (d *Database) AddEntry(e *Entry) {
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
