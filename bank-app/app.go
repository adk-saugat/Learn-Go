package main

import "fmt"

func main(){
	var balanceAmount float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
	fmt.Print("Your Choice: ")

	var choice int
	fmt.Scan(&choice)

	if choice == 1{
		fmt.Printf("\nYour balance is $%v.", balanceAmount)
	}else if choice == 2{
		fmt.Printf("\nYour deposit amount: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid Amount! Amout must be greater than 0.")
			return
		}

		balanceAmount += depositAmount
		fmt.Printf("\nBalance updated! New Amount: $%v\n", balanceAmount)
	}else if choice == 3{
		fmt.Printf("\nYour withdrawl amount: ")
		var withdrawlAmount float64
		fmt.Scan(&withdrawlAmount)

		if withdrawlAmount <= 0 {
			fmt.Println("Invalid Amount! Amout must be greater than 0.")
			return
		}

		if withdrawlAmount > balanceAmount {
			fmt.Println("You cannot withdraw more than what you have.")
			return
		}
			

		balanceAmount -= withdrawlAmount
		fmt.Printf("\nBalance updated! New Amount: $%v\n", balanceAmount)
	}else{

	}
}