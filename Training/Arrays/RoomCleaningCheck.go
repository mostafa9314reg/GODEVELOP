//In this Program we build a check cleaning Room  Service

package main

import "fmt"

type Room struct {
	name     string
	number   int
	cleanned bool
}

func RoomCheck(room [4]Room) {
	for i := 0; i < len(room); i++ {
		if room[i].cleanned {
			fmt.Printf("Room %s is cleared at %d.\n", room[i].name, room[i].number)
		} else {
			fmt.Printf("Room %s is not cleared at %d.\n", room[i].name, room[i].number)
		}
	}
}

func main() {
	rooms := [...]Room{{name: "Vip", number: 1}, {name: "Kitchen", number: 2}, {name: "Office", number: 3}, {name: "Girl", number: 4}}
	rooms[1].cleanned = true
	RoomCheck(rooms)
}
