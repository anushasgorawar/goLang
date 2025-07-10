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
		go calculate(i)
	}
	elapsed := time.Since(start)
	time.Sleep(2 * time.Second)
	fmt.Println("Total time taken: ", elapsed)
}
