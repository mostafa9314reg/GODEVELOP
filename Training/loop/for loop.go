//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	var result int = 0
	for i := 0; i < 50; i++ {
		result += i
		if result%3 == 0 && result%5 == 0 {
			fmt.Println("FizzBuzz The number is :", result)
		} else if result%5 == 0 {
			fmt.Println("Buzz The number is :", result)
		} else if result%3 == 0 {
			fmt.Println("Fizz The number is :", result)
		} else {
			fmt.Println(result)
		}

	}
}
