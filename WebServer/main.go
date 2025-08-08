package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
	fmt.Println("Endpoint Hit: homepage")
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to foo")
	fmt.Println("Endpoint Hit: foo")
}

func main() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/foo", foo)
	http.ListenAndServe("localhost:9090", nil) //nil->default mux
}
