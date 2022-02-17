package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Conference Go"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	//fmt.Printf("Welcome to our %v Booking Application \n", conferenceName)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The First Names of booking are %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end the program
				fmt.Printf("Our Conferencce is sold out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name you entered is INVALID")
			}
			if !isValidEmail {
				fmt.Println("Email you entered is INVALID")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of Tickets you entered is INVALID")
			}
		}
	}
}

func greetUsers(confName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our conference: %v", confName)
	fmt.Printf("We have a total of %v and still these many are remaining %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend here")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	//bookings := []string{}
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve the confirmation mail at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

}
