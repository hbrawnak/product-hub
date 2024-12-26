package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app App

func TestMain(m *testing.M) {
	err := app.Initialize(rootUser, rootPassword, "test")
	if err != nil {
		log.Fatal("Error initializing app")
	}

	createTable()
	m.Run()
}

func createTable() {
	query := `CREATE TABLE IF NOT EXISTS products (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    quantity int,
    price float(10,7),
    PRIMARY KEY (id)
	)`

	_, err := app.DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.DB.Exec("DELETE FROM products")
	app.DB.Exec("ALTER TABLE products AUTO_INCREMENT = 1")
}

func addProduct(name string, quantity int, price float64) {
	query := fmt.Sprintf("INSERT INTO products(name, quantity, price) VALUES ('%v', %v, %v);", name, quantity, price)
	_, err := app.DB.Exec(query)
	if err != nil {
		log.Println(err)
	}
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProduct("pens", 20, 5)
	request, _ := http.NewRequest("GET", "/api/products/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)
}

func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected status: %v, Received: %v", expectedStatusCode, actualStatusCode)
	}
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	app.Router.ServeHTTP(recorder, request)
	return recorder
}

func TestCreateProduct(t *testing.T) {
	clearTable()

	var product = []byte(`{"name": "Chair", "quantity": 5, "price": 10}`)
	req, _ := http.NewRequest("POST", "/api/products/", bytes.NewBuffer(product))
	req.Header.Set("Content-Type", "application/json")

	response := sendRequest(req)
	checkStatusCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "Chair" {
		t.Errorf("Expected name: %v, Received: %v", "Chair", m["name"])
	}

	if m["quantity"] != 5.0 {
		t.Errorf("Expected quantity: %v, Received: %v", 5.0, m["quantity"])
	}

	if m["price"] != 10.0 {
		t.Errorf("Expected price: %v, Received: %v", 10.0, m["price"])
	}
}
