package main

import "fmt"

func main() {
	var s string = "Hello"
	p := &s
	fmt.Println("Value of s is: ", s)
	fmt.Println("Address of s is: ", p)
	fmt.Println("Value of s is: ", *p)

	*p = "There"
	fmt.Println("Value of s is: ", s)
	fmt.Println("Address of s is: ", p)
	fmt.Println("Value of s is: ", *p)
}
