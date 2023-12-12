package main

import "fmt"

func main() {
	 conferenceName := "Go Conference"
	 const conferenceTickets = 50
	 remainingTickets := 50

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tocket here to attend!")
	
	var firstName string
	var lastName string
	var email string
	var userTickets int

	//Ask the user for his information
	fmt.Println("Enter your First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email Address")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets,email)

	fmt.Printf("%v tickets remaining for %v\n",remainingTickets, conferenceName)
}