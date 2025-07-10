package main

import "fmt"

func main() {
	for i := 8; i < 13; i++ {
		if a := 10; a < i {
			fmt.Println("b is greater")
		} else {
			fmt.Println("a is greater")
		}
	}
}
