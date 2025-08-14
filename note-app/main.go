package main

import (
	"errors"
	"fmt"
)

func main(){
	noteTitle, err := getUserInput("Enter note title: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	noteContent, err := getUserInput("Enter note content: ")
	if err != nil {
		fmt.Println(err)
		return
	}

}

func getUserInput(prompt string) (string, error){
	fmt.Println(prompt)

	var userInput string
	fmt.Scanln(&userInput)

	if userInput == ""{
		return "", errors.New("invalid input")
	}
	return userInput, nil
}	