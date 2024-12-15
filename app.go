package main

import "fmt"

func getFirstName() string {
	var firstname string
	fmt.Println("Please enter your first name")
	fmt.Scanln(&firstname)
	return firstname
}
