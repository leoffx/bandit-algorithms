package main

import (
	"fmt"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
	"github.com/leoffx/bandit-algorithms/lib/database"
	"github.com/leoffx/bandit-algorithms/lib/strategy"
)

const numRounds = 1000
const numArms = 10

func main() {
	bandit := bandit.NewBandit(numArms)
	db := database.NewDatabase()

	epsilon := 0.1
	strategy, err := strategy.NewRecoveringDifferenceSoftmax(db, epsilon)
	if err != nil {
		panic(err)
	}

	for i := 0; i < numRounds; i++ {
		eligibleArms := bandit.GetEligibleArms()
		chosenArm := strategy.ChooseArm(eligibleArms, db.Entries)
		reward := bandit.PullArm(chosenArm)
		db.Insert(&database.Entry{
			Round:        i,
			ChosenArm:    chosenArm,
			EligibleArms: eligibleArms,
			Reward:       reward,
		})
	}

	fmt.Println("Database: ", db)
}
