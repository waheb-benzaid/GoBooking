package main

import "fmt"

func main() {
	 conferenceName := "Go Conference"
	const conferenceTickets = 50
	 remainingTickets := 50

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tocket here to attend!")
	
	var userName string
	var userTickets int
	//Ask the user for his name
	
	userName="Tom"
	userTickets = 5

	fmt.Printf("User %v booked %v tickets \n", userName, userTickets)
}