package main

import "fmt"

func main() {
	// Name of variable [] array symbol then type of value
	// * * variable [](datatype)
	var fruits [2]string
	// assigning
	fruits[0] = "Orange"
	fruits[1] = "Banana"

	// * * Shorthand of declaring array
	veggies := []string{"Eggplant", "Squash", "Cabbage"}

	fmt.Println(fruits, veggies, len(veggies))
}
