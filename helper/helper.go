package helper

import "strings"

func ValidatUserInput(firstName string, lastName string, email string, userTickets int,remainingTickets int) (bool, bool, bool) {
	isValideName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValideName, isValidEmail, isValidTicketNumber
}