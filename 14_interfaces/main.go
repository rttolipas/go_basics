package main

import (
	"fmt"
	"math"
)

// ! Interfaces are methods that can be share or integrate within different Structs

// * *  delclare interface
type Shape interface {
	area() float64
}

// * * declare struct
type Circle struct {
	x, y, radius float64
}
type Square struct {
	w, h float64
}

// * * create method of a struct where we could implement interface
// ? * should we use area() or Shape?
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square) area() float64 {
	return s.w * s.h
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	square := Square{w: 5, h: 5}

	fmt.Printf("Circle Area %f\n", circle.area())
	fmt.Printf("Square Area %f\n", square.area())
}
