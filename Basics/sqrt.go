package main

import "fmt"

func sqrt(x float64) (y float64) {
	y = x / 2
	for i := 0; i < 10 && y*y > x; i++ {
		y -= ((y * y) - x) / (2 * y)
		fmt.Println(y)
	}
	return
}
func main() {
	n := 9.0
	fmt.Printf("Square root of %v is %v\n", n, sqrt(n))
}
