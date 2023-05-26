package main

import (
	"github.com/leoffx/a-sleeping-recovering-bandit-algorithm-for-optimizing-recurring-notifications/strategies"
	"github.com/leoffx/a-sleeping-recovering-bandit-algorithm-for-optimizing-recurring-notifications/structs"
)

const numRounds = 1000

func main() {
	history := structs.History{}

	epsilon := 0.1
	strategy, err := strategies.NewEpsilonGreedy(history, epsilon)

	if err != nil {
		panic(err)
	}

	bandit := structs.Bandit{}

	for i := 0; i < numRounds; i++ {
		arms := bandit.GetEligibleArms()
		arm := strategy.ChooseArm(arms)
		reward := arm.Draw()

	}

}
