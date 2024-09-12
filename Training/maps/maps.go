//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

///////////////////////////////////////////////////////

// I think if i use 2 maps its more effective and the code is much more minimal
// than the solution with using slice and maps so I tried first solution
package main

import (
	"fmt"
)

func displayServerStatus(mymap map[string]string) {
	//srvName := ""
	srvCount := 0
	srvStat := make(map[string]int)
	fmt.Println("This is All Servers :")
	for key, value := range mymap {
		fmt.Println(key, value)
		srvStat[value] += 1
		srvCount += 1

	}
	fmt.Println("This is All Servers count by each type :")
	for key, value := range srvStat {
		fmt.Println(key, value)
	}
}

//const (
//	Online      = 0
//	Offline     = 1
//	Maintenance = 2
//	Retired     = 3
//)

func main() {
	//servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	mymap := make(map[string]string)
	mymap["darkstar"] = "Online"
	mymap["aiur"] = "Online"
	mymap["omicron"] = "Online"
	mymap["w359"] = "Online"
	mymap["baseline"] = "Online"
	displayServerStatus(mymap)

}
