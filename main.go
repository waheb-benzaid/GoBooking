package main

import (
	"fmt"
	"go-booking/helper"
	"strconv"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings =make([]map[string]string, 0)

func main() {

	greetUsers()

	for{
		//Ask the user for his information
		firstName, lastName, email, userTickets:= getUserInput()
		
		isValideName, isValidEmail, isValidTicketNumber := helper.ValidatUserInput(firstName, lastName, email, userTickets,remainingTickets)

		if isValideName && isValidEmail && isValidTicketNumber {

			bookTicket (userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
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

func greetUsers()  {
	fmt.Println("Welcome to our users in the conference!")
	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tocket here to attend!")
}

func printFirstNames() [] string {
	
	firstNames :=[]string{}

	for _,booking := range bookings{
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
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

func bookTicket ( userTickets int, firstName string, lastName string, email string )  {
	remainingTickets = remainingTickets - userTickets
	
	//create a map fot a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatInt(int64(userTickets),10)

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets, conferenceName)
}