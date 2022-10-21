package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Package level variables
const conferenceName = "GoCon Conference"
const conferenceTickets uint = 50
const namesMinimumCharacters = 3
var availableTickets uint = 50
var bookings []string

func main() {
	welcome()

	for {
		var firstName string
		var lastName string
		var email string
		var boughtTickets uint

		firstName = captureFirstName(false)
		lastName = captureLastName(false)
		email = captureEmail(false)
		boughtTickets = captureBoughtTicket(false)

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

func captureBoughtTicket(validationFailed bool) uint {
	var boughtTickets uint
	if (validationFailed) {
		fmt.Printf("Maximum of %v ticket(s) can be bought\n", availableTickets)
		fmt.Printf("How many ticket(s) do you need? (%v available): ", availableTickets)
		fmt.Scan(&boughtTickets)
	} else {
		fmt.Printf("How many ticket(s) do you need? (%v available): ", availableTickets)
		fmt.Scan(&boughtTickets)
	}
	return validateBoughtTicket(boughtTickets)
}

func validateBoughtTicket(ticketCount uint) uint {
	isValidTicketNumber := ticketCount > 0 && ticketCount <= availableTickets

	for ; !isValidTicketNumber; {
		ticketCount = captureBoughtTicket(true)
		if ticketCount > 0 && ticketCount <= availableTickets {
			isValidTicketNumber = true
		}
	}
	return ticketCount
}

func captureEmail(validationFailed bool) string {
	var email string
	if (validationFailed) {
		fmt.Printf("Enter a valid Email Address: ")
		fmt.Scan(&email)
	} else {
		fmt.Print("Enter Email Address: ")
		fmt.Scan(&email)
	}

	return validateEmail(email)
}

func validateEmail(email string) string {
	isValidEmail := strings.Contains(email, "@")

	for ; !isValidEmail; {
		email = captureEmail(true)
		isValidEmail = strings.Contains(email, "@")
	}

	return email
}

func captureFirstName(validationFailed bool) string {
	var firstName string
	if (validationFailed) {
		fmt.Printf("Enter First Name (%v characters minimum): ", namesMinimumCharacters)
		fmt.Scan(&firstName)
	} else {
		fmt.Print("Enter First Name: ")
		fmt.Scan(&firstName)
	}
	return validateFirstName(firstName)
}

func validateFirstName(name string) string {
	isNumericFirstName := false
	if _, err := strconv.Atoi(name); err == nil {
		isNumericFirstName = true
	}

	isValidFirstName := len(name) >= namesMinimumCharacters && !isNumericFirstName

	for ; !isValidFirstName; {
		name = captureFirstName(true)

		if _, err := strconv.Atoi(name); err != nil {
			isNumericFirstName = false
		}
			
		isValidFirstName = len(name) >= namesMinimumCharacters && !isNumericFirstName
	}

	return name
}

func captureLastName(validationFailed bool) string {
	var lastName string
	if (validationFailed) {
		fmt.Printf("Enter Last Name (%v characters minimum): ", namesMinimumCharacters)
		fmt.Scan(&lastName)
	} else {
		fmt.Print("Enter Last Name: ")
		fmt.Scan(&lastName)
	}

	return validateLastName(lastName)
}

func validateLastName(name string) string {
	isNumericLastName := false
	if _, err := strconv.Atoi(name); err == nil {
		isNumericLastName = true
	}

	isValidLastName := len(name) >= namesMinimumCharacters && !isNumericLastName

	for ; !isValidLastName; {
		name = captureLastName(true)

		if _, err := strconv.Atoi(name); err != nil {
			isNumericLastName = false
		}

		isValidLastName = len(name) >= namesMinimumCharacters && !isNumericLastName
	}

	return name
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