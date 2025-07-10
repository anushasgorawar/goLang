package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("Anon Function")
		wg.Done()
	}()
	fmt.Println("Hello")
	wg.Wait()
}
