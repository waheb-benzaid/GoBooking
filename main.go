package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50
	var bookings [] string

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets, remainingTickets)
	fmt.Println("Get your tocket here to attend!")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	for{
		//Ask the user for his information
		fmt.Println("Enter your First Name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last Name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email Address")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets")
		fmt.Scan(&userTickets)
		
		isValideName := len(firstName) >= 2 && len(lastName) >=2
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValideName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("The Whole slice : %v \n",bookings)
			fmt.Printf("The first element slice : %v \n", bookings[0])
			fmt.Printf("The type of the slice : %T \n", bookings)
			fmt.Printf("The size of the slice : %v \n", len(bookings))
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets,email)
			fmt.Printf("%v tickets remaining for %v\n",remainingTickets, conferenceName)

			firstNames :=[]string{}

			for i := 0; i < len(bookings); i++ {
				var names = strings.Fields(bookings[i])
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of boking are: %v\n",firstNames)

			}else{
				if !isValideName {
					fmt.Println("Firstname or lastname is too short")
				}
				if !isValidEmail {
					fmt.Println("Email does not contain @")
				}
				if !isValidTicketNumber {
					fmt.Println("Number of tickets invalid!")
				}
			}
		}
}