//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func areaCalculate(rectangle Rectangle) int {
	return rectangle.Width * rectangle.Height
}

func perimeterCalculate(rectangle Rectangle) int {
	return (rectangle.Width + rectangle.Height) * 2
}

func main() {
	//area := areaCalculate(Rectangle{Width: 10, Height: 20})
	//perimeter := perimeterCalculate(Rectangle{Width: 10, Height: 20})
	data := Rectangle{
		Width:  0,
		Height: 0,
	}
	data.Height = 10
	data.Width = 20
	area := areaCalculate(data)
	perimeter := perimeterCalculate(data)
	fmt.Println(area)
	fmt.Println(perimeter)

}
