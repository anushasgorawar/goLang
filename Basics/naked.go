package main

import "fmt"

func nakedreturn(x int) (y, z int) {
	y = x
	z = y - x
	return
}

func main() {
	fmt.Println(nakedreturn(1))
}
