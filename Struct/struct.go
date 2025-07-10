package main

import "fmt"

type Circle struct {
	radius float64
	area   float64
}

func calculateArea(c *Circle) {
	area := 3.14 * c.radius * c.radius
	(*c).area = area
}

func main() {
	c := Circle{radius: 4, area: 0}
	fmt.Println(c)
	calculateArea(&c)
	fmt.Println(c)

}
