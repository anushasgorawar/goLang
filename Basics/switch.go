package main

import (
	"fmt"
	"runtime"
)

func main() {
	// os := runtime.GOOS
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("It's darwin.")
	case "linux":
		fmt.Println("It's linux.")
	default:
		fmt.Println("idk what OS this is.")
	}
}
