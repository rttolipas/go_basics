package main

import "fmt"

func main() {
	// * * Regular loop
	count := 10
	for i := 1; i <= count; i++ {
		fmt.Println(i)
	}

	// * * FizzBuzz
	for i := 0; i <= 30; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
