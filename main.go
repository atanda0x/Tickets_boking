package main

import (
	"booking/helper"
	"fmt"
	"sync"
	"time"
)

var availableTickets = 50000
var conferenceName = "Go Conference"

const conferenceTickets = 5000

var booking = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {

	greeter()

	firstName, lastName, email, userTickets := getUserDetails()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUser(firstName, lastName, email, userTickets, availableTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookingTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)

		firstNames := getFirstName()
		fmt.Printf("The first name of the bookings are: %v\n", firstNames)

		if availableTickets == 0 {
			fmt.Println("Our conference booking is booked, come back next year")
			// break
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
	wg.Wait()

}

func greeter() {
	fmt.Printf("Welcome to our %v booking application\nWe have total of %v tickets and %v are still available\nGet your tickets here to attend\n", conferenceName, conferenceTickets, availableTickets)
}

func getFirstName() []string {
	firstNames := []string{}
	for _, bookings := range booking {
		firstNames = append(firstNames, bookings.firstName)
	}
	return firstNames
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

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	booking = append(booking, userData)
	fmt.Printf("List of bookigs is %v\n", booking)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaininf for %v\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(15 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################################################################################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", tickets, email)
	fmt.Println("####################################################################################################")
	wg.Done()
}
