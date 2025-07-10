package main

import "fmt"

type Circle2 struct {
	radius float64
	area   float64
}

func (c *Circle2) calculateArea() {
	c.area = 3.14 * c.radius * c.radius
}

func main() {
	c := Circle2{radius: 4, area: 0}
	c.calculateArea()
	fmt.Println(c)
}
