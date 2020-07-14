package main

import "fmt"

// * * Type struct [Interface in TS]
// * * Note: Should be capital
type Person struct {
	name string
	age  int
}

// ! Method declaration (VALUE Reciever type of Method)
// * * (func) ("p" = Identifier | Person = Struct/Type)
// * * greet() = Method Name | void = return type
func (p Person) greet() {
	fmt.Printf("Hello my name is %v and I am %v years old", p.name, p.age)
}

// ! Method declaration (POINTER Reciever type of Method)
func (p *Person) incrementAge() {
	p.age++
}

// Method with params
func (p *Person) concatTitle(title string) {
	p.name = title + p.name
}

func main() {
	// * * Method initialization
	// * * Note: method could be access by passing Person struct to a variable
	person := Person{"Ralph Tolipas", 29}
	person.incrementAge()
	person.concatTitle("Atty.")
	person.greet()
}
