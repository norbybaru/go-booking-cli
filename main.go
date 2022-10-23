package main

import (
	"booking-cli/helper"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Package level variables
const conferenceName = "GoCon Conference"
const conferenceTickets uint = 50
type UserData struct {
	firstName string
	lastName string
	email string
	boughtTickets uint
}

// Empty slice of map
var bookings []UserData
var wg = sync.WaitGroup{}

func main() {
	welcome()

	input := helper.GetInputs()
	boughtTicket, _ := strconv.Atoi(input["boughtTicket"])

	user := UserData {
		firstName: input["firstName"],
		lastName: input["lastName"],
		email: input["email"],
		boughtTickets: uint(boughtTicket),
	}

	bookings = append(bookings, user)

	summary(user)
	// Increment or Sets number of WaitGroup counter used by goroutines to wait for
	wg.Add(1)
	go emailTickets(user.email, user.boughtTickets)

	if helper.AvailableTickets() == 0 {
		fmt.Printf("Our %v is currently sold out. Come back next year\n", conferenceName)
		//break
	}

	// Waits for WaitGroup counter to be decreased to 0
	wg.Wait()
}

func welcome() {
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, helper.AvailableTickets())
	fmt.Println("Get your tickets here to attend")
}

func summary(user UserData) {
	fmt.Printf("\nThank you %v %v for booking %v ticket(s).\n", user.firstName, user.lastName, user.boughtTickets)
	fmt.Printf("Booking confirmation will be sent to this email %v.\n", user.email)
	fmt.Printf("%v tickets is remaining\n", helper.AvailableTickets())

	firstNames := getFirstNamesFromBookings()
	fmt.Printf("These are all the bookings: %v\n\n", firstNames)
}

func getFirstNamesFromBookings() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func emailTickets(email string, boughtTickets uint) {
	time.Sleep(10 * time.Second)
	fmt.Println("#########")
	message := fmt.Sprintf("Sending %v ticket(s) to email %v\n", boughtTickets, email)
	fmt.Printf("Mail: %v", message)
	fmt.Println("#########")
	fmt.Printf("\n")
	// Decrement WaitGroup counter by 1
	wg.Done()
}