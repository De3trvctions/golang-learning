package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remaindingTickets uint = 50

	// fmt.Println("Welcome to", conferenceName, "application")
	// fmt.Println("Total of", conferenceTickets, "and", remaindingTickets, "are available")
	fmt.Printf("Welcome to %v application \n", conferenceName)
	fmt.Printf("Total of %v and %v are available \n", conferenceTickets, remaindingTickets)
	fmt.Println("Get your booking to attend")

	// print data type
	fmt.Printf("conferenceTickets is %T, remaintingTickets is %T, conferenceName is %T \n", conferenceTickets, remaindingTickets, conferenceName)

	// Declaring data type
	var userName string
	var userTickets int

	userName = "Alex"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets \n", userName, userTickets)

}
