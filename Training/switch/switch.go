//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func age(a int) int {
	return a
}
func main() {
	switch result := age(7); {
	case result == 0:

		fmt.Println("You are newborn")
	case result == 1 || result == 2 || result == 3:

		fmt.Println(" you are toddler")
	case 4 <= result && result <= 12:

		fmt.Println("you are a child")
	case result >= 13 && result <= 17:
		fmt.Println("you are teenager")
	case result >= 18:
		fmt.Println("you are adult")

	}

}
