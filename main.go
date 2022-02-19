package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Conference Go"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastname        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The First Names of booking are %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			//end the program
			fmt.Printf("Our Conferencce is sold out!")
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our conference: %v", conferenceName)
	fmt.Printf("We have a total of %v and still these many are remaining %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend here")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//bookings := []string{}
	// create a MAP for user
	var userData = UserData{
		firstName:       firstName,
		lastname:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve the confirmation mail at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	ticket := fmt.Sprintf("%v Ticket for %v %v", userTickets, firstName, lastName)
	fmt.Println("**************************")
	fmt.Printf("Sending ticket %v to email address: %v", ticket, email)
	fmt.Println("**************************")
	wg.Done()
}
