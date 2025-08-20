package main

import "fmt"


func main(){
	prices := []float64{10.99, 1.99}

	updatedPrices := append(prices, 2.99, 5.99)
	fmt.Println(updatedPrices)

	fmt.Println(prices)

	//removing a value
	prices = prices[1:]
	fmt.Println(prices)
}

// func main(){
// 	prices := [5]int{1,2,3,4}
// 	featuredPrices := prices[:3]
// 	fmt.Println(featuredPrices)
	
// 	// Slice is just a reference from memory & can change the original value
// 	featuredPrices[0] = 5
// 	fmt.Println(prices)
// 	fmt.Println(len(prices), cap(prices))
// }