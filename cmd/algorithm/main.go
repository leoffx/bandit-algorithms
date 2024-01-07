package main

import (
	"fmt"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
	"github.com/leoffx/bandit-algorithms/lib/database"
	"github.com/leoffx/bandit-algorithms/lib/strategy"
)

const numRounds = 100
const numArms = 3

func main() {
	bandit := bandit.NewBandit(numArms)
	db := database.NewDatabaseAggregator()

	strategy := strategy.NewEpsilonGreedy(0.3)

	for i := 0; i < numRounds; i++ {
		eligibleArms := bandit.GetEligibleArms()
		armToStats := db.ArmToStats()
		chosenArm := strategy.ChooseArm(eligibleArms, &armToStats)
		reward := bandit.PullArm(chosenArm)
		db.Insert(database.NewEntry(i, chosenArm, eligibleArms, reward))
	}

	fmt.Println("Database:\n", db)
	fmt.Println("Bandit:\n", bandit)
}
