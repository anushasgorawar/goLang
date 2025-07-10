package main

import "fmt"

var tr, fal bool

var (
	tru bool = false
)

func main() {
	var i int
	i = 10
	fmt.Println(i, tr, fal)

	var hel, thr string = "hello", "there"
	fmt.Println(hel, thr)

	var a, b, c = "int", "str", 100
	fmt.Println(a, b, c)

	d, e, f := "int", "str", 200
	fmt.Println(d, e, f)

	fmt.Printf("Type: %T and Value: %v", tru, tru)
}
