package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	for {

		// call getUserInput func
		firstName, lastName, email, userTickets := getUserInput()

		// call validateUserInput func
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// call bookTicket func here
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			// call firstNames function
			firstNames := getFirstNames()
			fmt.Printf("All the bookings %v\n", firstNames)

			if remainingTickets <= 0 {
				fmt.Println("Our conference is sold out. Sorry!")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email entered missing '@' sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets entered is invalid.")
			}
		}

	}
	wg.Wait()

}

func greetUser() {
	fmt.Printf("Welcome to our %v booking app! \n", conferenceName)
	fmt.Printf("We have %v tickets and %v still available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket now")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you'd like to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create a map for users
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of booking %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email on %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for the %v.\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}

/*
	HOW A SWITCH STATEMENT IS DONE
	team := "Manchester United"

	switch team {
		case "Manchester United":
			// some code for Manchester United is executed
		case "AC Milan":
			// some code execute
		case "Real Madrid":
			// some code executed
		case "Dortmund":
			// some code executed
		case "Sporting Lisbon":
			// some code executed
		default:
			fmt.Println("No one cares about the French league")
	}
*/
