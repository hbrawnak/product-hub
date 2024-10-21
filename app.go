package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize() error {
	connection := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DB_USER, DB_PASSWORD, DATABASE_NAME)
	var err error
	app.DB, err = sql.Open("mysql", connection)
	if err != nil {
		return err
	}

	app.Router = mux.NewRouter().StrictSlash(true)
	return nil
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products", getProducts).Methods("GET").Name("GetProducts")
}
