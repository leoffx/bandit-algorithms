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
	allArms := bandit.InitializeArms(numArms)
	history := database.NewHistory(allArms)

	epsilon := 0.1
	strategy, err := strategy.NewRecoveringDifferenceSoftmax(history, epsilon)
	if err != nil {
		panic(err)
	}

	for i := 0; i < numRounds; i++ {
		eligibleArms := bandit.GetEligibleArms(allArms)
		armToProbability := strategy.CalculateArmsProbabilities(eligibleArms)
		chosenArm := strategy.ChooseArm(armToProbability)
		reward := chosenArm.DrawReward()
		history.Update(chosenArm, armToProbability, reward)
	}

	fmt.Println("History: ", history)
}
