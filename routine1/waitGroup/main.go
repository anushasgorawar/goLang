package main

import (
	"fmt"
	"sync"
)

func calculate(i int, wg *sync.WaitGroup) {
	fmt.Println(i * i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go calculate(i, &wg)
	}
	wg.Wait()
	println("Last message")
}
