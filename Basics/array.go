package main

import "fmt"

func main() {
	fmt.Print("Hello")
	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr1, arr2)

	for ind, ele := range arr1 {
		fmt.Println("Value in ", ind, " index is: ", ele)
	}
	for ind, ele := range arr1 {
		defer fmt.Println("Value in ", ind, " index is: ", ele)
	}

	m := make(map[string]int)
	m["test"] = 100
	fmt.Println(m)
	fmt.Println(m["test"])
	delete(m, "test")
	fmt.Println(m)
}
