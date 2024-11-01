package main

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/", app.GetHome).Methods("GET").Name("home")
	app.Router.HandleFunc("/products", app.GetProductList).Methods("GET").Name("product-list")
	app.Router.HandleFunc("/products/{id}", app.GetProductById).Methods("GET").Name("product-details")
}
