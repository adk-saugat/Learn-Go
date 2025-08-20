package main

import "fmt"

func main(){
	prices := [5]int{1,2,3,4}
	featuredPrices := prices[:3]
	fmt.Println(featuredPrices)
	
	// Slice is just a reference from memory & can change the original value
	featuredPrices[0] = 5
	fmt.Println(prices)
	fmt.Println(len(prices), cap(prices))
}