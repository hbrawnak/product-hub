package main

import (
	"database/sql"
	"encoding/json"
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

func (app *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = product.createProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusCreated, product)
}

func (app *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])

	if err != nil {
		sendError(w, http.StatusInternalServerError, "Invalid product id")
		return
	}

	var product Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	product.ID = productId
	err = product.updateProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, product)
}

func (app *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])

	if err != nil {
		sendError(w, http.StatusInternalServerError, "Invalid product id")
		return
	}

	product := Product{ID: productId}

	err = product.DeleteProduct(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, map[string]string{"result": "success"})
}
