package main

import (
	"fmt"
	"strings"
)

func main()  {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	// greeting users
	greetUsers(conferenceName,conferenceTickets,remainingTickets) 

	var bookings [] string 
	
	for  {
		// Take user input
		firstName,lastName,email, userTickets := getUserInputs()

		// Validate user inputs
		isValidName,isValidEmail,isValidTicketNumber := validateUserInputs(firstName,lastName,email,userTickets,remainingTickets)

		if !isValidEmail || !isValidName || !isValidTicketNumber{
			fmt.Println("Your input is invalid pls correct.")
			continue
		}

		if ( userTickets > remainingTickets ){
			fmt.Printf("We only have %v tickets remaining, you cant book %v tickets\n",remainingTickets,userTickets)
			continue
		}

		bookings = append(bookings, firstName + " " + lastName)
		remainingTickets = remainingTickets - userTickets
	
		fmt.Printf("Hi, Thanks %v %v for booking %v tickets, you will receive confirmation mail on your mail id: %v soon\n",
		firstName,lastName,userTickets,email)

		// get firstNames of users
		firstNames := getFirstNames(bookings)
	
		fmt.Printf("First name of all bookings: %v \n",firstNames)
		fmt.Printf("Tickets remaining: %v\n", remainingTickets)
		
		if remainingTickets == 0 {
			fmt.Println("Our conference booking is booked out, come again next year.")
			break
		}
	}

}

func greetUsers(conferenceName string,conferenceTickets uint,remainingTickets uint)  {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}


func getUserInputs() (string,string,string,uint)  {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Take user details:
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName,lastName,email, userTickets
}

func validateUserInputs(firstName string,lastName string,email string,userTickets uint,remainingTickets uint) (bool,bool,bool) {
	
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumber :=  userTickets > 0 && userTickets <= remainingTickets 
	
	return isValidName, isValidEmail, isValidTicketNumber
}

func getFirstNames(bookings [] string) [] string {
	firstNames := [] string {}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames,names[0])
	}

	return firstNames 
}