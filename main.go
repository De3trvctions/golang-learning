package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remaindingTickets = 50

	fmt.Println("Welcome to", conferenceName, "application")
	fmt.Println("Total of", conferenceTickets, "and", remaindingTickets, "are available")
	fmt.Println("Get your booking to attend")

}
