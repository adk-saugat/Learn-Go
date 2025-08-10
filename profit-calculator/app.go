package main

import (
	"errors"
	"fmt"
	"os"
)

func main(){
	revenue, err := getUserInput("Revenue")
	if err != nil{
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses")
	if err != nil{
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate")
	if err != nil{
		fmt.Println(err)
		return
	}
	earningBeforeTax, profit, ratio := calculateValues(revenue, expenses, taxRate)

	writeToFile(earningBeforeTax, profit, ratio)


	fmt.Println(earningBeforeTax)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func writeToFile(ebt, profit, ratio float64) {
	dataString := fmt.Sprintf("EBT:%.1f\nProfit:%.1f\nRatio:%.2f", ebt, profit, ratio)
	os.WriteFile("data.txt", []byte(dataString), 0644)
}

func calculateValues(revenue, expenses, taxRate float64) (float64 ,float64, float64){
	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax * (1 - taxRate / 100)
	ratio := earningBeforeTax / profit
	return earningBeforeTax, profit, ratio
}

func getUserInput(infoText string) (float64, error){
	var userInput float64
	fmt.Printf("%v: ", infoText)
	fmt.Scan(&userInput)

	if userInput <= 0{
		return 0, errors.New("input should be greater than 0")
	}
	return userInput, nil
}