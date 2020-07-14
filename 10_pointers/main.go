package main

import "fmt"

func main() {
	// * * Pointer is assigning value from memory location
	a := "hello"
	b := &a
	fmt.Println(b)

	// * *  Getting value from Pointer use (*) asterisk

	fmt.Println(*b)
	fmt.Println(*&a)

	// * * Change the orinal value of a variable from an assign pointer
	*b = "World"
	fmt.Println(a) // Output will be "World"
}
