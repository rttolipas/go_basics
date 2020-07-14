package main

import "fmt"

func main() {
	// * * Declare map kv
	emails := map[string]string{
		"Bob":  "bob@email.com",
		"John": "john@email.com",
	}

	for key, value := range emails {
		fmt.Printf("%v has the email of %v\n", key, value)
	}

}
