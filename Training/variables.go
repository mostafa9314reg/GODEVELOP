// Variables Asignment checking
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

	str1, str2 := "String1", "String2"
	str3, str2 := "String3", "String2"
	fmt.Println(str1, str2, str3)

	var (
		y     = 7
		z int = 9
	)
	fmt.Println("sum of y+z is :", y+z)

	word1, word2, _ := "Hello", "World", "Mostafa"
	fmt.Println(word1, word2)
}
