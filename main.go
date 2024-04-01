package main

import "fmt"

func main()  {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50


	fmt.Printf("Type of conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n",conferenceName,conferenceTickets,remainingTickets)
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Ram"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets \n", userName,userTickets)

}
