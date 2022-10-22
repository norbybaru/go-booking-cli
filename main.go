package main

import (
	"fmt"
	"strings"
	"booking-cli/helper"
)

// Package level variables
const conferenceName = "GoCon Conference"
const conferenceTickets uint = 50
var availableTickets uint = helper.AvailableTickets()
var bookings []string

func main() {
	welcome()

	for {
		firstName, lastName, email, boughtTickets := helper.GetInputs()

		bookings = append(bookings, firstName + " " + lastName)
		availableTickets = availableTickets - uint(boughtTickets)

		summary(firstName, lastName, email, boughtTickets)

		if availableTickets == 0 {
			fmt.Printf("Our %v is currently sold out. Come back next year\n", conferenceName)
			break
		}
	}
}

func welcome() {
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")
}

func summary(firstName string, lastName string, email string, boughtTickets uint) {
	fmt.Printf("\nThank you %v %v for booking %v ticket(s).\n", firstName, lastName, boughtTickets)
	fmt.Printf("Booking confirmation will be sent to this email %v.\n", email)
	fmt.Printf("%v tickets is remaining\n", availableTickets)

	firstNames := getFirstNamesFromBookings()
	fmt.Printf("These are all the bookings: %v\n\n", firstNames)
}

func getFirstNamesFromBookings() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstName := strings.Fields(booking)[0]
		firstNames = append(firstNames, firstName)
	}

	return firstNames
}