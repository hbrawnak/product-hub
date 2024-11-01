package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (app *App) GetHome(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Welcome to the Product Hub API!"}
	sendResponse(w, http.StatusOK, response)
}

func (app *App) GetProductList(w http.ResponseWriter, r *http.Request) {
	products, err := getProducts(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, products)
}

func (app *App) GetProductById(w http.ResponseWriter, r *http.Request) {
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
