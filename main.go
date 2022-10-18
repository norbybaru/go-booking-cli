package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName = "GoCon Conference"
	const conferenceTickets int = 50
	var availableTickets uint = 50
	var bookings []string

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var boughtTickets int

		fmt.Println("Enter First Name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter Last Name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter Email Address:")
		fmt.Scan(&email)

		fmt.Println("How many ticket(s) do you need?:")
		fmt.Scan(&boughtTickets)

		bookings = append(bookings, firstName + " " + lastName)
		availableTickets = availableTickets - uint(boughtTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			firstName := strings.Fields(booking)[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Printf("Thank you %v %v for booking %v tickets. Booking confirmation will be sent to this email %v\n", firstName, lastName, boughtTickets, email)
		fmt.Printf("%v tickets is remaining\n", availableTickets)
		fmt.Printf("These are all the bookings: %v\n", firstNames)
	}
}