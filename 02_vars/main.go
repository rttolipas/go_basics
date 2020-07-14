package main

import "fmt"

func main() {
	// * * No need to explicitly define type when it is a "String" type
	// const name string = "Ralph" // hey, you can use "const" here
	var age uint8 = 29

	// * * Shorthand declaration, Note: it will auto detect the type
	// name := "Ralph"

	// * * Multiple single declaration
	name, email := "Ralph", "ralph@email.com"

	// * * Display output
	fmt.Println("Hi my name is ", name, ". I am ", age, " years old ", email)
	fmt.Printf("%T", age)
}
