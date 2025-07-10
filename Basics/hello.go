package main

import "fmt"

func main() {
	var last string
	var age int
	fmt.Println("Hello world")
	name := "Anusha"
	fmt.Printf("%v %T", name, name)
	fmt.Print("What is your last name and age? ")
	fmt.Scanf("%s %d", &last, &age)
	fmt.Print("Welcome, Anusha ", last, ". You are ", age)
	fmt.Println(name, last, age)

	var n = "n"
	m := "n"
	print(n, m)
}
