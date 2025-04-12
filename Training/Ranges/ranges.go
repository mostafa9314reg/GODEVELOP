// Here we are trying to use ranges for our loop

package main

import "fmt"

func main() {
	slice := []string{"my", "name", "is", "Mostafa"}

	for i, element := range slice {
		fmt.Println(i, element)
		for _, element2 := range element {
			fmt.Printf(" %q\n ", element2)
		}
	}

}
