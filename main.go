package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	remainingTickets = -1

	fmt.Printf("conferenceTicket is %T, remaining tickets is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to \"%v\" booking application\n", conferenceName)
	fmt.Printf("We have  total of: %v tickets and still tickets: %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTicket int
	// ask user for their name

	userName = "Tom"
	userTicket = 2
	fmt.Printf("User %v booked %v tickets.", userName, userTicket)
}
