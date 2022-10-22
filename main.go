package main

import (
	"booking-cli/helper"
	"fmt"
	"strconv"
)

// Package level variables
const conferenceName = "GoCon Conference"
const conferenceTickets uint = 50
var availableTickets uint = helper.AvailableTickets()
// Empty slice of map
var bookings []map[string]string

func main() {
	welcome()

	for {
		userData := helper.GetInputs()
		bookings = append(bookings, userData)
	
		boughtTicket, _ := strconv.Atoi(userData["boughtTicket"])
		availableTickets = availableTickets - uint(boughtTicket)

		summary(userData["firstName"], userData["lastName"], userData["email"], userData["boughtTicket"])

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

func summary(firstName string, lastName string, email string, boughtTickets string) {
	fmt.Printf("\nThank you %v %v for booking %v ticket(s).\n", firstName, lastName, boughtTickets)
	fmt.Printf("Booking confirmation will be sent to this email %v.\n", email)
	fmt.Printf("%v tickets is remaining\n", availableTickets)

	firstNames := getFirstNamesFromBookings()
	fmt.Printf("These are all the bookings: %v\n\n", firstNames)
}

func getFirstNamesFromBookings() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}