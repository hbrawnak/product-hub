package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize() error {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DbUser, DbPassword, DatabaseName)
	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes()
	return nil
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {
	sendResponse(w, statusCode, map[string]string{"error": err})
}

func (app *App) getHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page")
	response := map[string]string{"message": "Welcome to the Product Hub API!"}
	sendResponse(w, http.StatusOK, response)
}

func (app *App) getProductList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Product list")
	products, err := getProducts(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, products)
}

func (app *App) getProductById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Product by id")
	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])

	if err != nil {
		sendError(w, http.StatusInternalServerError, "Invalid product id")
		return
	}

	product := Product{ID: productId}
	err = product.GetProduct(app.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			sendError(w, http.StatusNotFound, "Product not found")
		default:
			sendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	sendResponse(w, http.StatusOK, product)
}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/", app.getHome).Methods("GET").Name("home")
	app.Router.HandleFunc("/products", app.getProductList).Methods("GET").Name("home")
	app.Router.HandleFunc("/products/{id}", app.getProductById).Methods("GET").Name("home")
}
