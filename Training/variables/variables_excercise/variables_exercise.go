//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var favColor = "Red"
	birthYear, age := 1990, 34

	var (
		firstInitial = 2024
		lastInitial  = 2024
	)
	var daysOfAge int
	daysOfAge = age * 365

	fmt.Println("my favorite color is : ", favColor)
	fmt.Println("my age is : ", age)
	fmt.Println("my birth year is : ", birthYear)
	fmt.Println("my first initial is : ", firstInitial)
	fmt.Println("my last initial is : ", lastInitial)
	fmt.Println("my daysOfAge is : ", daysOfAge)
}
