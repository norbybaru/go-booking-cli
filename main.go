package main

import (
	"fmt"
)

func main() {
	const conferenceName = "GoCon Conference"
	const conferenceTickets = 50
	var availableTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var boughtTickets int

	userName = "John Doe"
	boughtTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, boughtTickets)
}