package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Data struct {
	id   int
	name string
}

func main() {
	connnectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1)/%v", USER, PASSWORD, DATABASE)
	db, err := sql.Open("mysql", connnectionString)
	defer db.Close()

	printError(err)

	rows, err := db.Query("SELECT * FROM data")
	printError(err)

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		printError(err)
		fmt.Println(data)
	}
}

func printError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
