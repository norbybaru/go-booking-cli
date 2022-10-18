package main

import "fmt"

func main() {
	const conferenceName = "GoCon Conference"
	const conferenceTickets = 50
	var availableTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")
}