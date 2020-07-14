package main

import "fmt"

func main() {
	// * * Define map
	// map[type]type

	// emails := make(map[string]string)

	// emails["Bob"] = "bob@email.com"
	// emails["John"] = "john@email.com"

	// * * Declare map kv
	emails := map[string]string{
		"Bob":  "bob@email.com",
		"John": "john@email.com",
	}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete map (array, key)
	delete(emails, "Bob")
	fmt.Println(emails)
}
