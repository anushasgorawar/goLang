package main

import "fmt"

type Circle struct {
	radius float64
}

// Interface has all method definitions
type Shape interface {
	CalculateArea() float64
}

// Struct Method declaration
func (c Circle) CalculateArea() float64 {
	return c.radius * c.radius * 3.14
}

// Generic Function for Shape interface
func CalculateArea(s Shape) float64 {
	return s.CalculateArea()
}

func main() {
	c := Circle{radius: 5}
	fmt.Printf("Area of cicle c is: %v", c.CalculateArea())
	fmt.Printf("Area of cicle c is: %v", CalculateArea(c))
}
