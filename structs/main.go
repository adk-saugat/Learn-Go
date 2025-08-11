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

func (appUser *user )outputUserData(){
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthdate)
}

func (appuser *user) clearUserName(){
	appuser.firstName = ""
	appuser.lastName = ""
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

	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
