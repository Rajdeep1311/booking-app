package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to our %v Booking Application\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "and", remainingTickets, "are left")
	fmt.Println("Get your tickets now")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// asks users for their name
	fmt.Printf("Enter your frist name : ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets to be booked : ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining Tickets for the %v", remainingTickets, conferenceName)
}
