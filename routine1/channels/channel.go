package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("UNBUFFERED CHANNEL")
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go buy(ch, &wg)
	go sell(ch, &wg)
	wg.Wait()

	fmt.Println("BUFFERED CHANNEL")
	buff := make(chan int, 2)
	buff <- 10
	buff <- 12
	fmt.Println(<-buff)
	fmt.Println(<-buff)
	// fmt.Println(<-buff) Will give deadlock
	buff <- 14
	fmt.Println(<-buff)
	buff <- 16
	val, ok := <-buff
	fmt.Println(val, ok)
	close(buff) // if not closed, next line will give error
	val, ok = <-buff
	fmt.Println(val, ok)

}
func buy(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Bought 10 apples")
	ch <- 10
	wg.Done()
}
func sell(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Selling 10 apples")
	fmt.Println(<-ch)
	fmt.Println("Sold 10 apples")
	wg.Done()
}
