package main

import (
	"bookinghub/utility"
	"bookinghub/utils"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Bookinghub")
	utils.Sum()
	ulfile := utility.Name()
	var userdata = make(map[string]string)
	userdata["akash"] = "shivank"
	userdata["aka"] = "shivank"
	userdata["ak"] = "shivank"

	fmt.Println(userdata)
	for key, value := range userdata {
		fmt.Println("Key:", key, "Value:", value)
	}

	fmt.Println(ulfile)
	// fmt.Println(mulfile)
}

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Welcome to Bookinghub")
// 	var totalTickets = 50
// 	var remainingTickets = totalTickets
// 	var bookingarr []string
// 	mulfile := name()
// 	fmt.Println(mulfile)
// 	for remainingTickets > 0 {
// 		firstname := getFirstName()
// 		lastname := getLastName()
// 		isValid := validateName(firstname, lastname)
// 		if !isValid {
// 			fmt.Println("Please enter the length of first name and last name more than 2 and try again")
// 			continue
// 		}

// 		email := getEmail()
// 		isValidEmail := validateEmail(email)
// 		if !isValidEmail {
// 			fmt.Println("Hello, please enter your correct email address and try again")
// 			break
// 		}

// 		usertickets := getTickets()
// 		if usertickets > remainingTickets {
// 			fmt.Printf("Hello sir, we have only %v tickets remaining and you need more than that. Please try again sir \n", remainingTickets)
// 			continue
// 		}

// 		if len(bookingarr) == 0 {
// 			bookingarr, remainingTickets = bookTickets(firstname, lastname, usertickets, remainingTickets)
// 		} else {
// 			bookingarr = append(bookingarr, firstname+" "+lastname)
// 			remainingTickets = remainingTickets - usertickets
// 		}

// 		fmt.Printf("Total %v tickets are left\n ", remainingTickets)
// 		displayFirstNames(bookingarr)
// 	}
// }

// func main() {

// fmt.Println("Welcome! to the booking website")
// fmt.Println("You will get your ticket here.")

// var totalTickets int
// var remainingTickets int

// // Ask for total tickets and read the value
// fmt.Println("Please enter your total tickets:")
// fmt.Scanf("%d", &totalTickets)

// // Ask for remaining tickets and read the value
// fmt.Println("Please enter your remaining tickets:")
// fmt.Scanf("%d", &remainingTickets)

// // Print the total and remaining tickets
// fmt.Printf("Total number of tickets: %d and total number of remaining tickets: %d\n", totalTickets, remainingTickets)

// var username string
// var useremail string
// var usertickets int

// // Reading username
// fmt.Printf("Enter username:\n")
// fmt.Scanln(&username) // Use Scanln for strings to handle newlines correctly

// // Reading email address
// fmt.Printf("Enter the email address:\n")
// fmt.Scanln(&useremail) // Use Scanln for string
// fmt.Printf("please enter the ticket you have required")
// fmt.Scanf("%d", &usertickets)
// remainingTickets = remainingTickets - usertickets

// fmt.Printf("totalnumber of remaining ticket is now %d", remainingTickets)

// var bookingarray [50]int
// bookingarray[0] = 10
// booksarr := []string{}
// booksarr = append(booksarr, "akash")
// fmt.Println(bookingarray)

// func main() {

// 	var totoalticket = 50
// 	var remainingTickets = 50
// 	bookingarr := []string{}

// 	fmt.Printf("Total number of tickets: %d\n", totoalticket)
// 	fmt.Println("Hello sir, please enter your details.")

// 	for {

// 		var name string
// 		fmt.Println("Hi, please enter your name:")
// 		fmt.Scanln(&name) // Reads a single input for the name

// 		var email string
// 		fmt.Println("Please enter your email:")
// 		fmt.Scanln(&email) // Reads a single input for the email
// 		if remainingTickets == 0 {
// 			fmt.Println(name, "sorry we have not tickets now ")
// 			break
// 		}
// 		var usertickets int
// 		fmt.Println("Please enter how many tickets you need:")
// 		fmt.Scanf("%d\n", &usertickets) // Reads a number for the ticket count

// 		// Update remaining tickets
// 		remainingTickets = remainingTickets - usertickets
// 		fmt.Printf("Total number of remaining tickets: %d\n", remainingTickets)

// 		// Append the name to the booking array
// 		bookingarr = append(bookingarr, name+" "+email)
// 		fmt.Println("Current bookings:", bookingarr)
// 		firstnames := []string{}

// 		for _, bookingarr := range bookingarr {
// 			var names = strings.Fields(bookingarr)
// 			var firstname = names[0]
// 			firstnames = append(firstnames, firstname)

// 		}
// 		lastnames := []string{}
// 		for _, bookingarr := range bookingarr {

// 			var names = strings.Fields(bookingarr)
// 			var lastname = names[1]
// 			lastnames = append(lastnames, lastname)

// 		}
// 		fmt.Println("lastname are", lastnames)

// 		fmt.Println("firstname are", firstnames)

// 	}
// }

// func main() {
// 	fmt.Println("Welcome to Bookinghub")
// 	var totoalticket = 50
// 	bookingarr := []string{}
// 	var remainingTickets = totoalticket
// 	for remainingTickets > 0 {
// 		var firstname string
// 		fmt.Println("please enter you firstname")
// 		fmt.Scanln(&firstname)
// 		var lastname string
// 		fmt.Print("please enter yourlastname\n")
// 		fmt.Scanln(&lastname)
// 		isValid := len(firstname) > 2 && len(lastname) > 2
// 		if !isValid {
// 			fmt.Println("Please enter the length of first name and last name more than 2 and try again")
// 			continue
// 		}
// 		var email string
// 		fmt.Println("please enter your email")
// 		fmt.Scanln(&email)
// 		isvalidemail := strings.Contains(email, "@")
// 		if !isvalidemail {
// 			fmt.Println("hello please enter your correct mail address and try again")
// 			break
// 		}
// 		var usertickets int
// 		fmt.Println("hello please enter how many tickets you want")
// 		fmt.Scanln(&usertickets)
// 		if usertickets > remainingTickets {
// 			fmt.Printf("hello sir we have only %v ticket are remaining and you need more then that please try again sir \n", remainingTickets)
// 			continue
// 		}
// 		bookingarr = append(bookingarr, firstname+" "+lastname)
// 		remainingTickets = remainingTickets - usertickets
// 		fmt.Printf("total %v of ticket are left\n ", remainingTickets)
// 		firstnames := []string{}
// 		for _, bookingarr := range bookingarr {
// 			name := strings.Fields(bookingarr)
// 			firstnames = append(firstnames, name[0])

// 		}
// 		fmt.Println("first name array is ", firstnames)
// 		// fmt.Println("bookig array is ", bookingarr)

// 	}

// }
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// Function to get user's first name
