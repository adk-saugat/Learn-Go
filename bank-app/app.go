package main

import (
	"fmt"
	"os"
	"strconv"
)

func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func getBalanceFromFile() float64{
	data, _ := os.ReadFile("balance.txt")
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance

}

func main(){
	var balanceAmount float64 = getBalanceFromFile()

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Print("Your Choice: ")
	
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("\nYour balance is $%v.", balanceAmount)
		case 2:
			fmt.Printf("\nYour deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Println("Invalid Amount! Amout must be greater than 0.")
				continue
			}
	
			balanceAmount += depositAmount
			fmt.Printf("\nBalance updated! New Amount: $%v\n", balanceAmount)
			writeBalanceToFile(balanceAmount)
		case 3:
			fmt.Printf("\nYour withdrawl amount: ")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)
	
			if withdrawlAmount <= 0 {
				fmt.Println("Invalid Amount! Amout must be greater than 0.")
				continue
			}
	
			if withdrawlAmount > balanceAmount {
				fmt.Println("You cannot withdraw more than what you have.")
				continue
			}
				
	
			balanceAmount -= withdrawlAmount
			fmt.Printf("\nBalance updated! New Amount: $%v\n", balanceAmount)
			writeBalanceToFile(balanceAmount)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
	
}