package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
}

func getProducts(db *sql.DB) ([]Product, error) {
	query := `SELECT id, name, quantity, price FROM products`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	products := []Product{}
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (p *Product) GetProduct(db *sql.DB) error {
	query := fmt.Sprintf("SELECT name, quantity, price FROM products WHERE id = %v", p.ID)
	row := db.QueryRow(query)
	err := row.Scan(&p.Name, &p.Quantity, &p.Price)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) CreateProduct(db *sql.DB) error {
	query := fmt.Sprintf("insert into products(name, quantity, price) values('%v', %v, %v)", p.Name, p.Quantity, p.Price)
	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = int(id)
	return nil
}

func (p *Product) UpdateProduct(db *sql.DB) error {
	query := fmt.Sprintf("update products set name = '%v', quantity = %v, price = %v where id = %v", p.Name, p.Quantity, p.Price, p.ID)
	result, err := db.Exec(query)
	rowAffected, err := result.RowsAffected()
	if rowAffected == 0 {
		return errors.New("product does not exist")
	}

	return err
}

func (p *Product) DeleteProduct(db *sql.DB) error {
	query := fmt.Sprintf("delete from products where id = %v", p.ID)
	result, err := db.Exec(query)
	rowAffected, err := result.RowsAffected()
	if rowAffected == 0 {
		return errors.New("product does not exist")
	}

	return err
}
