package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName            string
	lastName             string
	emailAddress         string
	numberOfTickets      uint
	isOptedForNewsLetter bool
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	firstName, lastName, emailAddress, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, emailAddress)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, emailAddress)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are %v\n", firstNames)

		var noTicketsRemaining bool = remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("Our conference is booked out. Come back next year")
		}
	} else {
		fmt.Println("Your input data is invalid. Try again")
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
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
	var emailAddress string
	var userTickets uint

	fmt.Println("Enter your fitst name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&emailAddress)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAddress, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		emailAddress:    emailAddress,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, emailAddress)
	fmt.Println("##################")
	wg.Done()
}
