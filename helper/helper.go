package helper

import "strings"

func ValidateUser(firstName string, lastName string, email string, userTickets int, availableTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 3 && len(lastName) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= availableTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
