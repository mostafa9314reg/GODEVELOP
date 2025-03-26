// Here we are trying to  make an application which print Routes

package main

import "fmt"

func iteration(slice []string) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

}

func main() {
	Route := []string{"R1", "R2", "R3", "R4"}
	fmt.Println("All Routes Are:")
	fmt.Println(Route)
	fmt.Println("Route R5 Added")
	Route = append(Route, "R5")
	fmt.Println(Route)
	fmt.Println("Route R6 Added")
	Route = append(Route, "R6")
	fmt.Println(Route)
}
