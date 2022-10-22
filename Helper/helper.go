package helper

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
)

const namesMinimumCharacters = 3
const availableTickets uint = 50

func AvailableTickets() uint {
	return availableTickets
}

func GetInputs() (string, string, string, uint) {
	firstName := captureFirstName(false)
	lastName := captureLastName(false)
	email := captureEmail(false)
	boughtTicket := captureBoughtTicket(false)

	return firstName, lastName, email, boughtTicket
}

func CaptureInput(field string) (string, error) {
	switch field {
		case "firstName":
			return captureFirstName(false), nil
		case "lastName":
			return captureLastName(false), nil
		case "email":
			return captureEmail(false), nil
		case "boughtTicket":
			return strconv.FormatUint(uint64(captureBoughtTicket(false)), 10), nil
		default:
			fmt.Println("Invalid field")
			return "", errors.New("invalid field: " + field)
	}
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
	return validateInputNames(firstName, "firstName")
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

	return validateInputNames(lastName, "lastName")
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

func validateEmail(email string) string {
	isValidEmail := strings.Contains(email, "@")

	for ; !isValidEmail; {
		email = captureEmail(true)
		isValidEmail = strings.Contains(email, "@")
	}

	return email
}

func validateInputNames(value string, input string) string {
	isNumericName := false
	if _, err := strconv.Atoi(value); err == nil {
		isNumericName = true
	}

	isValidName := len(value) >= namesMinimumCharacters && !isNumericName

	for ; !isValidName; {
		switch input {
			case "firstName":
				value = captureFirstName(true)
			case "lastName":
				value = captureLastName(true)
			default:
				fmt.Printf("Invalid input value: %v", input)
		}

		if _, err := strconv.Atoi(value); err != nil {
			isNumericName = false
		}
			
		isValidName = len(value) >= namesMinimumCharacters && !isNumericName
	}

	return value
}