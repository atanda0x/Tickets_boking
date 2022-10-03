package main

import (
	"fmt"
	"strings"
)

var availableTickets = 50000
var conferenceName = "Go Conference"

const conferenceTickets = 50000

var booking = []string{}

func main() {

	greeter()

	for {

		firstName, lastName, email, userTickets := getUserDetails()

		isValidName, isValidEmail, isValidTicketNumber := validateUser(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookingTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstName()
			fmt.Printf("The first name of the bookings are: %v\n", firstNames)

			if availableTickets == 0 {
				fmt.Println("Our conference booking is booked, come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name you entered is short")
			} else if !isValidEmail {
				fmt.Println("Email doesn't have @ ")
			} else if !isValidTicketNumber {
				fmt.Println("Number of ticket is invalid")
			}
		}
	}

}

func greeter() {
	fmt.Printf("Welcome to our %v booking application\nWe have total of %v tickets and %v are still available\nGet your tickets here to attend\n", conferenceName, conferenceTickets, availableTickets)
}

func getFirstName() []string {
	firstNames := []string{}
	for _, bookings := range booking {
		name := strings.Fields(bookings)
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}

func validateUser(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 3 && len(lastName) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= availableTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserDetails() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookingTickets(userTickets int, firstName string, lastName string, email string) {
	remainingTickets := availableTickets - userTickets
	booking = append(booking, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaininf for %v\n", remainingTickets, conferenceName)
}
