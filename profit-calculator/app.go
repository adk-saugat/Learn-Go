package main

import "fmt"

func main(){
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax * (1 - taxRate / 100)
	ratio := earningBeforeTax / profit

	fmt.Println(earningBeforeTax)
	fmt.Println(profit)
	fmt.Println(ratio)
}