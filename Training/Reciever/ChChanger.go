//Here we making a program to change Video Game character changer

package main

import (
	"fmt"
)

type Player struct {
	name      string
	Health    int
	MaxHealth int
	Energy    int
	MaxEnergy int
	MinEnergy int
	MinHealth int
}

func (player *Player) Healing(amount int) {
	player.Health += amount
	if player.Health > player.MaxHealth {
		player.Health = player.MaxHealth
	}
	fmt.Println("Healing Player by ", amount)
}

func (player *Player) ConsumeEnergy(amount int) {
	player.Energy -= amount
	if player.Energy < player.MinEnergy {
		player.Energy = 0
		fmt.Println("Energy of Player Just Finished And He Died !!!!!!!! Sorry !!!!!!!")
	}
}

func (player *Player) Damage(amount int) {
	player.Health -= amount
	if player.Health < player.MinHealth {
		player.Health = 0
		fmt.Println("Health of Player Just Finished And He Died !!!!!!!! Sorry !!!!!!!")
	}
}

//func(player *Player)PlayerStats() {
//	fmt.Println("These Are Player Stats:")
//	fmt.Println(player)
//}

func main() {

	player := Player{
		name:      "Mostafa",
		Health:    100,
		MaxHealth: 100,
		Energy:    100,
		MaxEnergy: 100,
		MinEnergy: 1,
		MinHealth: 1,
	}
	fmt.Println("lets consume energy by 10")
	player.ConsumeEnergy(10)
	fmt.Println("Now player Stats is :")
	fmt.Println(player)

}
