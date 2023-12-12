package main

import "strings"

func validatUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	isValideName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValideName, isValidEmail, isValidTicketNumber
}