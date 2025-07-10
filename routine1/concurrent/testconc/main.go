package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)
	go start()
	fmt.Println("Main function")
}

func start() {
	fmt.Println("start function")
	go process()
}

func process() {
	fmt.Println("process function")
}
