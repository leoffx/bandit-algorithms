package main

import (
	"fmt"

	"github.com/leoffx/a-sleeping-recovering-bandit-algorithm-for-optimizing-recurring-notifications/strategies"
	"github.com/leoffx/a-sleeping-recovering-bandit-algorithm-for-optimizing-recurring-notifications/structs"
)

const numRounds = 1000
const numArms = 10

func main() {
	history := structs.History{ArmToStats: make(map[*structs.Arm]structs.Stats)}
	allArms := structs.InitializeArms(numArms)

	epsilon := 0.1
	strategy, err := strategies.NewEpsilonGreedy(history, epsilon)

	if err != nil {
		panic(err)
	}

	for i := 0; i < numRounds; i++ {
		eligibleArms := structs.GetEligibleArms(allArms)
		arm := strategy.ChooseArm(eligibleArms)
		reward := arm.DrawReward()
		history.Update(arm, reward)
	}

	fmt.Println("History: ", history)
}
