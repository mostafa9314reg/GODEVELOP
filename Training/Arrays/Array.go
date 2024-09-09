//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func ProductList(products []Product) {
	//for _, product := range products {
	//	fmt.Println(product.Name)
	//	fmt.Println(product.Price)
	//}
	items := 0
	cost := 0.0
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].Name)
		fmt.Println(products[i].Price)
		items += 1
		cost += products[i].Price
		if i == len(products)-1 {
			fmt.Println("The last item is", products[i].Name)
		}

	}
	fmt.Println("Total items are:", items)
	fmt.Println("Total cost is:", cost)
}
func main() {
	product := [...]Product{
		{Name: "Shirt"},
		{Name: "Shoes"},
		{Name: "Beer"},
		{Name: "Book"},
	}
	fmt.Println("Before Assigning Price")
	ProductList(product[:])
	fmt.Println("After Assigning Price")
	product[0].Price = 100
	product[1].Price = 150
	product[2].Price = 20
	product[3].Price = 30
	//product[4].Price = 40
	ProductList(product[:])

}
