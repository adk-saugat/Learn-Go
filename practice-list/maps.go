package main

import "fmt"

func main() {
	capitalCities := map[string]string{}

	capitalCities["Nepal"] = "Kathmandu"
	capitalCities["United States"] = "Washington D.C"

	for key, value := range capitalCities {
		fmt.Println(key, "->", value)
	}

	delete(capitalCities, "United States")

	fmt.Println(capitalCities)

	
}
