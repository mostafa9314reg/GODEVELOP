//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
)

func randGenerator(n int) int {
	// Generate a random Int type number between 1 and 10
	//using this formula "rand.Intn(max - min + 1) + min"
	randNum := rand.Intn(n-1+1) + 1
	return randNum
}
func rollCount(roll int, side int) int {
	summRollResult := 0
	for i := 1; i <= roll; i++ {
		randResult := randGenerator(side)
		fmt.Println(" in roll: ", i)
		fmt.Println("dice Shows:", randResult)
		summRollResult += randResult
	}
	fmt.Println("SummRollResult:", summRollResult)
	return summRollResult
}

func diceCount(roll int, side int, count int) {
	sumOfDices := 0
	for i := 1; i <= count; i++ {
		sumOfDices = sumOfDices + rollCount(roll, side)
	}
	fmt.Println("sum of all dices rolls is :", sumOfDices)
}
func main() {
	roll := 2
	side := 3
	dice := 3
	if roll == 2 && dice == 2 {
		fmt.Println("Snake eyes")
	} else if roll == 7 {
		fmt.Println("Lucky 7")
	} else if roll%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
	diceCount(roll, side, dice)

}
