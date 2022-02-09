package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Conference Go"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to our %v Booking Application \n", conferenceName)
	fmt.Printf("We have a total of %v and still these many are remaining %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend here")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		fmt.Println("Enter your Firstname")
		fmt.Scan(&firstName)
		fmt.Println("Enter your Lastname")
		fmt.Scan(&lastName)
		fmt.Println("Enter your Email")
		fmt.Scan(&email)
		fmt.Println("Enter number of Tickets required")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("we only have %v tickets remaining, you can't buy %v tickets \n", remainingTickets, userTickets)
			continue
		}
		remainingTickets = remainingTickets - userTickets

		bookings := []string{}
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve the confirmation mail at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets are remaining for %v\n", conferenceTickets, remainingTickets)
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The First names of bookings %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			//end the program
			fmt.Printf("Our Conferencce is sold out!")
			break
		}
	}
	//var userName string

}
