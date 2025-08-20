package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	// Products
	products := []Product{{
		title: "Coke",
		id: 1,
		price: 2.69,
	}, {
		title: "Lays",
		id: 2,
		price: 3.29,
	}}

	products = append(products, Product{
		title: "Gatorade",
		id: 3,
		price: 3.59,
	})
	fmt.Println(products)

	// Hobbies
	hobbies := [3]string{"Code", "Read", "Play Games"}

	fmt.Println(hobbies)

	firstHobby := hobbies[:1]
	fmt.Println("First hobby:", firstHobby)
	fmt.Println("Second & Third hobby:", hobbies[1:])

	first2Hobbies := firstHobby[:2]
	fmt.Println("\nFirst two hobbies:", first2Hobbies)

	last2Hobbies := first2Hobbies[1:3]
	fmt.Println("Last two hobbies:", last2Hobbies)

	// CourseGoal
	courseGoals := []string{"Learn Go", "Get Interview"}

	courseGoals[1] = "Do leetcode"
	courseGoals = append(courseGoals, "Get lean")
	fmt.Println(courseGoals)


}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.