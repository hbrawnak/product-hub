package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Product struct {
	Id       string
	Name     string
	Quantity float64
	Price    float64
}

var Products []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
}

func getProductList(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: productList")
	err := json.NewEncoder(w).Encode(Products)
	if err != nil {
		return
	}
}

func getProductDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["productId"]
	for _, Product := range Products {
		if string(Product.Id) == key {
			json.NewEncoder(w).Encode(Product)
		}
	}
	log.Println(key)
}

func handleRequests() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/products/", getProductList)
	muxRouter.HandleFunc("/products/{productId}", getProductDetail)
	muxRouter.HandleFunc("/", homePage)
	err := http.ListenAndServe(":10000", muxRouter)
	if err != nil {
		return
	}
}

func main() {
	Products = []Product{
		Product{Id: "1", Name: "Chair", Quantity: 100, Price: 100.00},
		Product{Id: "2", Name: "Desk", Quantity: 150, Price: 200.00},
	}

	handleRequests()
}
