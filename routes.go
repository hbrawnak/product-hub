package main

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/", app.GetHome).Methods("GET").Name("home")

	//Product API
	products := app.Router.PathPrefix("/api/products").Subrouter()
	products.HandleFunc("/", app.GetProductList).Methods("GET").Name("product-list")
	products.HandleFunc("/{id}", app.GetProductById).Methods("GET").Name("product-details")
}
