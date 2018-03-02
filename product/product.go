package product

import (
	"database/sql"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Repository interface {
	getProduct(db *sql.DB, id int) (Product, error)
	updateProduct(db *sql.DB, product Product) error
	deleteProduct(db *sql.DB, id int) error
	createProduct(db *sql.DB, product Product) (int, error)
	getProducts(db *sql.DB, start, count int) ([]Product, error)
}
