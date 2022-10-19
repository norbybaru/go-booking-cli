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
		var boughtTickets uint

		fmt.Print("Enter First Name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter Last Name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter Email Address: ")
		fmt.Scan(&email)

		fmt.Printf("How many ticket(s) do you need? (%v available): ", availableTickets)
		fmt.Scan(&boughtTickets)

		isValidTicketNumber := boughtTickets > 0 && boughtTickets <= availableTickets

		for ; !isValidTicketNumber; {
			fmt.Printf("Maximum of %v ticket(s) can be bought\n", availableTickets)
			fmt.Printf("How many ticket(s) do you need? (%v available): ", availableTickets)
			fmt.Scan(&boughtTickets)
			if boughtTickets > 0 && boughtTickets <= availableTickets {
				isValidTicketNumber = true
			}
		}

		bookings = append(bookings, firstName + " " + lastName)
		availableTickets = availableTickets - uint(boughtTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			firstName := strings.Fields(booking)[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Printf("\nThank you %v %v for booking %v tickets.\n", firstName, lastName, boughtTickets)
		fmt.Printf("Booking confirmation will be sent to this email %v.\n", email)
		fmt.Printf("%v tickets is remaining\n", availableTickets)
		fmt.Printf("These are all the bookings: %v\n\n", firstNames)

		if availableTickets == 0 {
			fmt.Printf("Our %v is currently sold out. Come back next year\n", conferenceName)
			break
		}
	}
}