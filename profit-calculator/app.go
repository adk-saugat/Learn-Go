package main

import "fmt"

func main(){
	revenue := getUserInput("Revenue")
	expenses := getUserInput("Expenses")
	taxRate := getUserInput("Tax Rate")

	earningBeforeTax, profit, ratio := calculateValues(revenue, expenses, taxRate)

	fmt.Println(earningBeforeTax)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func calculateValues(revenue, expenses, taxRate float64) (float64 ,float64, float64){
	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax * (1 - taxRate / 100)
	ratio := earningBeforeTax / profit
	return earningBeforeTax, profit, ratio
}

func getUserInput(infoText string) float64{
	var userInput float64
	fmt.Printf("%v: ", infoText)
	fmt.Scan(&userInput)
	return userInput
}