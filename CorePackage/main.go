package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello World"
	fmt.Println(text)

	check := "World"
	fmt.Println(strings.Contains(text, check))
	fmt.Println(strings.Contains(check, text))

	fmt.Println(strings.ReplaceAll(text, "l", "y"))

	fmt.Println(strings.Count(text, "l"))
}
