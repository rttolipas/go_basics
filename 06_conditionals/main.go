package main

import "fmt"

func main() {
	// * * IF Else
	x := 10
	y := 5
	if x > y {
		fmt.Printf("%v is greater the %v\n", y, x)
	}

	// * * Switch
	color := "red"
	switch color {
	case "red":
		fmt.Println("color is red")
	default:
		fmt.Println("color is not red")
	}
}
