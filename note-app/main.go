package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/note/note"
)

func main(){
	noteTitle := getUserInput("Enter note title: ")	
	noteContent := getUserInput("Enter note content: ")

	userNote, err := note.New(noteTitle, noteContent)

	if err != nil {
		fmt.Println(err)
		return 
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed!")
		return		
	}

	fmt.Println("Saving the note succeeded!")
}

func getUserInput(prompt string) (string){
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil{
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	
	return text
}	