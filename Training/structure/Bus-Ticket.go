// Here we are trying to make a Bus Ticket application
//some passengers have tickets that show their seats and been boarded and Bus that have these Passengers.

package main

import (
	"fmt"
)

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	Name string
	Seat Passenger
}

func main() {
	mostafa := Passenger{"Mostafa", 23, true}
	fmt.Println(mostafa)
	var (
		Alli    = Passenger{"Alli", 22, true}
		Hossein = Passenger{"Hossein", 20, true}
	)
	fmt.Println("we have Three member Mostafa Ali And Hossein :", mostafa, Alli, Hossein)
	fmt.Println("Hossein Ticket number is", Hossein.TicketNumber)

	Alli.Boarded = false
	bus := Bus{"MaralVip", mostafa}
	bus2 := Bus{"Scania", Hossein}
	fmt.Println("Bus1 and Bus2:", bus, bus2)
}
