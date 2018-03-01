package product

import (
	"database/sql"
)

type Product struct {
	Name  string
	Price float64
}

type Repository interface {
	getProduct(db *sql.DB) (Product, error)
	updateProduct(db *sql.DB) error
	deleteProduct(db *sql.DB) error
	createProduct(db *sql.DB, product Product) (int, error)
	getProducts(db *sql.DB, start, count int) ([]Product, error)
}
