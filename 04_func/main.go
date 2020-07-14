package main

import "fmt"

func greetings(name string) string {
	return "Hello " + name
}

// * * If wll declare same, this will work
func sum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greetings("Ralph"))
	fmt.Println(sum(1, 2))
}
