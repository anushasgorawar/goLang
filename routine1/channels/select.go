package main

import (
	"fmt"
	"sync"
)

func goOne(ch1 chan string, wg *sync.WaitGroup) {
	ch1 <- "Channel-1"
	wg.Done()
}
func goTwo(ch2 chan string, wg *sync.WaitGroup) {
	ch2 <- "Channel-2"
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1, &wg)
	go goTwo(ch2, &wg)

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	}
}
