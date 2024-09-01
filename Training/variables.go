package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	//fmt.Println("Hello World Test First")

	var a = 1
	var b, e = 2, 5
	c := 3
	var d int = 4
	result := a + b + c + d + e
	fmt.Println("The result of add is  :", result)

	var x int
	var str string
	fmt.Println("Default value int is:", x, "Default value string is:", str)
}
