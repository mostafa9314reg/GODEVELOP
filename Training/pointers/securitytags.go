//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func Activate(tag *SecurityTag) {
	*tag = Active
}

func Deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func Checkout(items []Item) {
	fmt.Println("Checking....")
	for i := 0; i < len(items); i++ {
		Deactivate(&items[i].tag)
	}
}

func main() {
	Book := Item{name: "Book", tag: Active}
	Shirt := Item{name: "Shirt", tag: Active}
	Laptop := Item{name: "Laptop", tag: Active}

	items := []Item{Book, Shirt, Laptop}

	fmt.Println(items)
	//Deactivate(&items[0].tag)
	Checkout(items)
	fmt.Println(items)
	Deactivate(&items[0].tag)
	fmt.Println(items)
}
