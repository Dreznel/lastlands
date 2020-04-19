package roller

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RollDice(amount, sides int) []int {
	min := 1
	max := sides

	var results []int
	for i := 0; i < amount; i++ {
		value := rand.Intn(max-min+1) + min
		results = append(results, value)
	}

	return results
}

func RollAbilityScore() int {
	rolls := RollDice(3, 6)
	fmt.Println(fmt.Sprintf("You rolled %d, %d, and %d", rolls[0], rolls[1], rolls[2]))
	lowestRoll, err := min(rolls)
	if err != nil {
		fmt.Println("problem occured rolling ability score. Returning value of -99")
		return -99
	}
	return lowestRoll

}

func min(values []int) (int, error) {
	if len(values) == 0 {
		return -1, errors.New("no values specified")
	}

	minValue := values[0]
	for _, value := range values {
		if value < minValue {
			minValue = value
		}
	}

	return minValue, nil
}

func max(values []int) (int, error) {
	if len(values) == 0 {
		return -1, errors.New("no values specified")
	}

	maxValue := values[0]
	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue, nil
}
