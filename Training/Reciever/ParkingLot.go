package main

import "fmt"

type Space struct {
	occupied bool
}

type Spaces struct {
	spaces []Space
}

func OccupySpace(lot *Spaces, spacenumber int) {
	lot.spaces[spacenumber-1].occupied = true
}
func (lot *Spaces) OccupySpace(spacenumber int) {
	lot.spaces[spacenumber-1].occupied = true
}

func (lot *Spaces) Vaccate(spacenumber int) {
	lot.spaces[spacenumber-1].occupied = false
}

func main() {

	lot := Spaces{spaces: make([]Space, 4)}

	fmt.Println("Before Occupying")
	fmt.Println(lot)
	fmt.Println("After Occupying 1,2")
	OccupySpace(&lot, 1)
	lot.OccupySpace(2)
	fmt.Println(lot)
	fmt.Println("After Vacating 1")
	lot.Vaccate(1)
	fmt.Println(lot)

}
