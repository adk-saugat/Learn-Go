package main

import "fmt"

func main(){
	age := 32 // regular variable
	agePointer := &age // Pointer

	fmt.Println("Age:", *agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Printf("Age (adult years): %v",age)
}

func editAgeToAdultYears(age *int){
	*age = *age - 18
}
