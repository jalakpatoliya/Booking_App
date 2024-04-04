package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings [] UserData

type UserData struct {
	firstName string
	lastName string
	email string
	noOfTickets uint
}

var wg = sync.WaitGroup{}

func main()  {

	// greeting users
	greetUsers() 

	// Take user input
	firstName,lastName,email, userTickets := getUserInputs()

	// Validate user inputs
	isValidName,isValidEmail,isValidTicketNumber := validateUserInputs(firstName,lastName,email,userTickets)

	if isValidEmail && isValidName && isValidTicketNumber {

		// book tickets
		bookTickets(firstName,lastName,email, userTickets)
		
		wg.Add(1)
		// send tickets on new thread using go
		go sendTickets(userTickets,firstName,lastName,email )

		// get firstNames of users
		firstNames := getFirstNames()
	
		fmt.Printf("First name of all bookings: %v \n",firstNames)
		fmt.Printf("Tickets remaining: %v\n", remainingTickets)
		fmt.Printf("Bookings: %v\n",bookings)
		
		if remainingTickets == 0 {
			fmt.Println("Our conference booking is booked out, come again next year.")
			// break
		}

	} else {
		if !isValidEmail || !isValidName || !isValidTicketNumber{
			fmt.Println("Your input is invalid pls correct.")
			// continue
		}

		if ( userTickets > remainingTickets ){
			fmt.Printf("We only have %v tickets remaining, you cant book %v tickets\n",remainingTickets,userTickets)
			// continue
		}
	}

	wg.Wait()
}

func greetUsers()  {
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

func getFirstNames() [] string {
	firstNames := [] string {}

	for _, booking := range bookings {
		firstNames = append(firstNames,booking.firstName)
	}

	return firstNames 
}

func bookTickets(firstName string,lastName string,email string,userTickets uint)  {
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		noOfTickets: userTickets,
	}

	// bookings = append(bookings, firstName + " " + lastName)
	bookings = append(bookings, userData)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Hi, Thanks %v %v for booking %v tickets, you will receive confirmation mail on your mail id: %v soon\n",
	firstName,lastName,userTickets,email)
}

func sendTickets(userTickets uint,firstName string, lastName string, email string)  {
	// simulating time delay
	time.Sleep(10*time.Second)

	var tickets = fmt.Sprintf("%v tickets for %v %v",userTickets,firstName,lastName)
	fmt.Println("################")
	fmt.Printf("Sending %v to email %v\n", tickets,email)
	fmt.Println("################")

	wg.Done()
}