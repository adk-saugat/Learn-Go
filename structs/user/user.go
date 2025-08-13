package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName string
	Birthdate string
	CreatedAt time.Time
}

func New(firstName, lastName, birthdate string) (*User, error){

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("values should not be empty") 
	}
	return &User{
		FirstName: firstName,
		LastName: lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}, nil
}

func (appUser *User) OutputUserData() {
	fmt.Println(appUser.FirstName, appUser.LastName, appUser.Birthdate)
}

func (appUser *User) ClearUserName() {
	appUser.FirstName = ""
	appUser.LastName = ""
	appUser.Birthdate=""
}
