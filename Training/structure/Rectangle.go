// Here we try to make an application to calculate area and primeter of a Trctangle

package main

import "fmt"

type Coordinates struct {
	x int
	y int
}

type Rectangle struct {
	a, b Coordinates
}

func width(rect Rectangle) int {
	width := rect.b.x - rect.a.x
	return width
}

func length(rect Rectangle) int {
	len := rect.b.y - rect.a.y
	return len
}

func Area(rect Rectangle) int {
	area := width(rect) * length(rect)
	return area

}

func Perimeter(rect Rectangle) int {
	perimeter := (width(rect) + length(rect)) * 2
	return perimeter
}

func main() {

	rectangle := Rectangle{a: Coordinates{10, 20}, b: Coordinates{20, 30}}
	area := Area(rectangle)
	primeter := Perimeter(rectangle)
	fmt.Println("This is the Area of Rectangle:", area)
	fmt.Println("This is  the Priemeter of Rectangle:", primeter)
}
