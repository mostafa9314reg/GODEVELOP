// Here we are making a grocery shopping list and doing some functions on it

package main

import "fmt"

func main() {

	shoppinglist := make(map[string]int)
	shoppinglist["Bread"] = 1
	shoppinglist["Milk"] = 2
	shoppinglist["Jam"] = 1
	shoppinglist["Egg"] = 3
	fmt.Println("Here is the initial ShoppingList :", shoppinglist)
	delete(shoppinglist, "Milk")
	fmt.Println("just delete Milk :", shoppinglist)
	shoppinglist["Bread"] += 1
	fmt.Println("added 1 bread", shoppinglist)
	cereal, found := shoppinglist["Cereal"]
	fmt.Println("Have cereal ?  ")
	if found {
		fmt.Println("found cereal", cereal)
	} else {
		fmt.Println("no cereal found")
	}
	totalItems := 0

	for _, value := range shoppinglist {
		totalItems += value
	}

	fmt.Println("total items", totalItems)
}
