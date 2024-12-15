package main

import (
	"fmt"
	"strings"
)

func name() string {
	return "akash"
}

// Function to get user's last name
func getLastName() string {
	var lastname string
	fmt.Print("Please enter your last name\n")
	fmt.Scanln(&lastname)
	return lastname
}

// Function to validate user's name
func validateName(firstname, lastname string) bool {
	return len(firstname) > 2 && len(lastname) > 2
}

// Function to get user's email
func getEmail() string {
	var email string
	fmt.Println("Please enter your email")
	fmt.Scanln(&email)
	return email
}

// Function to validate user's email
func validateEmail(email string) bool {
	return strings.Contains(email, "@")
}

// Function to get number of tickets from user
func getTickets() int {
	var usertickets int
	fmt.Println("Hello, please enter how many tickets you want")
	fmt.Scanln(&usertickets)
	return usertickets
}

// Function to book tickets
func bookTickets(firstname, lastname string, usertickets int, remainingTickets int) ([]string, int) {
	bookingarr := []string{firstname + " " + lastname}
	remainingTickets = remainingTickets - usertickets
	return bookingarr, remainingTickets
}

// Function to display first names of booked users
func displayFirstNames(bookingarr []string) {
	firstnames := []string{}
	for _, booking := range bookingarr {
		name := strings.Fields(booking)
		firstnames = append(firstnames, name[0])
	}
	fmt.Println("First name array is ", firstnames)
}
