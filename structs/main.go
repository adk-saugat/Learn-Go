package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	var appUser user
	appUser = user{
		firstName: userFirstName,
		lastName:userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	outputUserData(&appUser)
}

func outputUserData(appUser *user){
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
