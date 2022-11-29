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
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Declaring array required an array size
	var bookings []string

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remaindingTickets = remaindingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole array %v \n", bookings)
	fmt.Printf("The first value %v \n", bookings[0])
	fmt.Printf("Slice type: %T \n", bookings)
	fmt.Printf("Slice length %v \n", len(bookings))

	fmt.Printf("Thankiu %v %vfor booking %v ticket. You will receive confirmation at %v.\n", firstName, lastName, email, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remaindingTickets, conferenceName)
}
