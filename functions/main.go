package main

import "fmt"

type CalcType func(int, int) int

func main(){
	result := calculateValue(1, 3, add)
	fmt.Println("Result:", result)

	result2 := calculateValue(1, 3, subtract)
	fmt.Println("Result:", result2)
}

func calculateValue(first int, second int, calcType CalcType) int{
	return calcType(first, second)
}

func add(first, second int) int{
	return first + second
}

func subtract(first, second int) int{
	return first - second
}