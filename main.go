package main

import "fmt"

func main()  {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50


	fmt.Printf("Type of conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n",conferenceName,conferenceTickets,remainingTickets)
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	
	// Take user details:
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Hi, Thanks %v %v for booking %v tickets, you will receive confirmation mail on your mail id: %v soon\n",
	 firstName,lastName,userTickets,email)
}
