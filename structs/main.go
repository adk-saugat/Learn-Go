package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthdate string) (*user, error){

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("Values should not be empty!") 
	}
	return &user{
		firstName: firstName,
		lastName:lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
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
	// var appUser *user
	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
