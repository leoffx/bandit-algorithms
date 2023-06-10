package main

import (
	"fmt"

	"github.com/leoffx/a-sleeping-recovering-bandit-algorithm-for-optimizing-recurring-notifications/strategies"
	"github.com/leoffx/a-sleeping-recovering-bandit-algorithm-for-optimizing-recurring-notifications/structs"
)

const numRounds = 1000
const numArms = 10

func main() {
	allArms := structs.InitializeArms(numArms)
	history := structs.NewHistory(allArms)

	epsilon := 0.1
	strategy, err := strategies.NewEpsilonGreedy(history, epsilon)
	if err != nil {
		panic(err)
	}

	for i := 0; i < numRounds; i++ {
		eligibleArms := structs.GetEligibleArms(allArms)
		armToProbability := strategy.CalculateArmsProbabilities(eligibleArms)
		chosenArm := strategy.ChooseArm(armToProbability)
		reward := chosenArm.DrawReward()
		history.Update(chosenArm, armToProbability, reward)
	}

	fmt.Println("History: ", history)
}
