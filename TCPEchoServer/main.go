package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// go run main.go 9000
func main() {
	if len(os.Args) < 2 {
		// fmt.Println("Port Number Missing")
		log.Fatal("Missing Port Number")
		os.Exit(1)
	}
	// fmt.Printf("Request to %v\n", os.Args[1])
	port := fmt.Sprintf(":%v", os.Args[1])

	listener, err := net.Listen("tcp", port)
	// creates a socket to the network interface and port number
	if err != nil {
		log.Fatal("Failed to create a listener: ", err)
		//2025/07/26 14:17:36 Failed to create a listener: listen tcp :80: bind: address already in use
		os.Exit(1)
	}
	defer listener.Close() //cleans the socket in the end
	// fmt.Println(listener)  //Request to 9000&{0xc00010e000 {<nil> 0 {false 0 0 0} 0}}

	fmt.Println("Listening on", listener.Addr())

	for {
		conn, err := listener.Accept() //Conn is an interface
		//Accept function blocks and waits till a client connects
		if err != nil {
			log.Fatal("Failed to accept requets: ", err)
			continue //we want the for loop running
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// fmt.Println(conn) //&{{0xc0000be080}}
	reader := bufio.NewReader(conn)
	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("request: %s", bytes)
		line := fmt.Sprintf("Echo: %s", bytes)

		fmt.Printf("response: %s", line)

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("failed to write data, err:", err)
			return
		}
	}
}
