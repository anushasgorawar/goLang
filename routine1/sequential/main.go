package main

import (
	"fmt"
	"time"
)

func calculate(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		calculate(i)
	}
	elapsed := time.Since(start)
	fmt.Println("Total time taken: ", elapsed)
}
