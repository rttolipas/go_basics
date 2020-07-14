package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
}

func main() {
	person := Person{name: "Ralph Tolipas", age: 29, gender: "M"}
	fmt.Println(person)

	// Other way of declaration
	person2 := Person{"James Bond", 28, "M"}
	person2.age = 30
	fmt.Println(person2)

}
