package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func addint(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Sum of two numbers is %g", add(5, 6))
	fmt.Println("Sum of two numbers is %g", addint(6, 6))
}
