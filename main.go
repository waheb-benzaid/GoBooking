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

	greetUsers(conferenceName,conferenceTickets,remainingTickets)

	for{
		//Ask the user for his information
		firstName, lastName, email, userTickets:= getUserInput()
		
		isValideName, isValidEmail, isValidTicketNumber := validatUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValideName && isValidEmail && isValidTicketNumber {

			bookTicket (remainingTickets, userTickets, bookings, firstName, lastName, conferenceName, email)

			firstNames:=printFirstNames(bookings)
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


func greetUsers(conferenceName string, confTickets int, remainingTickets int)  {
	fmt.Println("Welcome to our users in the conference!")
	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",confTickets, remainingTickets)
	fmt.Println("Get your tocket here to attend!")
}

func printFirstNames(bookings[] string) [] string {
	
	firstNames :=[]string{}

	for i := 0; i < len(bookings); i++ {
		var names = strings.Fields(bookings[i])
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validatUserInput(firstName string,lastName string,email string,userTickets int,remainingTickets int) (bool,bool,bool) {
	
	isValideName := len(firstName) >= 2 && len(lastName) >=2
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValideName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int
	
	fmt.Println("Enter your First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email Address")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket (remainingTickets int, userTickets int, bookings []string, firstName string, lastName string, conferenceName string, email string )  {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets, conferenceName)

}