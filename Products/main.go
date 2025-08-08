package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id       string
	Name     string
	Price    float64
	Quantity int
}

var products []Product //initialising a slice

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnAllProducts")
	json.NewEncoder(w).Encode(products) //NewEncoder writes to w, Encode function writes the json encoding to the stream. We are converting our data to json and writing it to ResponseWriter stream
}

func returnOneProduct(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	path := r.URL.Path[len("/product/"):] //Not an elegant way to fetch the id
	for _, product := range products {
		if product.Id == path {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
	log.Println("Endpoint hit: Homepage")
}

func handleRequests() {
	http.HandleFunc("/", homepage) //register routes
	http.HandleFunc("/products", returnAllProducts)
	http.HandleFunc("/product/", returnOneProduct)
	http.ListenAndServe("localhost:9090", nil)
}

func main() {
	products = []Product{
		{
			Id:       "1",
			Name:     "Shoes",
			Price:    79.99,
			Quantity: 2,
		}, {
			Id:       "2",
			Name:     "Socks",
			Price:    4.99,
			Quantity: 2,
		},
	}
	handleRequests()
}
