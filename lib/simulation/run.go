package simulation

import (
	"fmt"

	"github.com/leoffx/bandit-algorithms/lib/bandit"
	"github.com/leoffx/bandit-algorithms/lib/database"
	"github.com/leoffx/bandit-algorithms/lib/strategy"
)

func Run(numArms int, numRounds int, strategy strategy.Strategy) (*bandit.Bandit, *database.DatabaseAggregator) {
	bandit := bandit.NewBandit(numArms)
	db := database.NewDatabaseAggregator()

	for i := 0; i < numRounds; i++ {
		eligibleArms := bandit.GetEligibleArms()
		armToStats := db.ArmToStats()
		armToScore := strategy.ScoreArms(eligibleArms, armToStats)
		chosenArm := strategy.ChooseArm(armToScore)

		reward := bandit.PullArm(chosenArm)
		db.AddEntry(database.NewEntry(i, chosenArm, armToScore, reward))
	}

	// fmt.Println("Database:\n", db)
	fmt.Println("Bandit:\n", bandit)
	return bandit, db
}
