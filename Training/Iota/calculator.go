//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

const (
	Add = iota
	Sub = iota
	Mul = iota
	Div = iota
)

type oparation int

func (op oparation) calculate(lhs, rhs float64) float64 {
	switch op {
	case Add:
		return lhs + rhs
	case Sub:
		return lhs - rhs
	case Mul:
		return lhs * rhs

	case Div:
		return lhs / rhs

	}
	panic("unhandled operation")

}

func main() {

	add := oparation(Add)
	fmt.Println(add.calculate(2, 2)) // = 4

	sub := oparation(Sub)
	fmt.Println(sub.calculate(10, 3)) // = 7

	mul := oparation(Mul)
	fmt.Println(mul.calculate(3, 3)) // = 9

	div := oparation(Div)
	fmt.Println(div.calculate(100, 2)) // = 50

}
