// --Summary:
//
//	Use functions to perform basic operations and print some
//	information to the terminal.
//
// --Requirements:
//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
//   - Write a function that returns any message, and call it from within
//     fmt.Println()
//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
//   - Write a function that returns any number
//   - Write a function that returns any two numbers
//   - Add three numbers together using any combination of the existing functions.
//   - Print the result
//   - Call every function at least once
package main

import (
	"fmt"
	_ "fmt"
	"math/rand"
	//"unsafe"
)

func greeting(name string) {
	fmt.Println("Greeting ", name)
}

func messager() string {
	message := "Hello from messager"
	return message
}

func add(a, b, c int) int {
	return a + b + c
}

func random() int {
	return rand.Int()

}
func numreturn() (int, int) {
	a := 5
	b := 10
	return a, b
}

func main() {
	greeting("Mostafa")
	fmt.Println(messager())
	fmt.Println("sum of 2,3,5 is :", add(2, 3, 5))
	fmt.Println("The random number is :", random())
	a, b := numreturn()
	fmt.Println("two number returned are :", a, b)
	d := random()
	fmt.Println("result of adding three numbers od three difrent functionis  :", add(a, b, d))

}
