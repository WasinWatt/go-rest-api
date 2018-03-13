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
	GetProduct(db *sql.DB, id int) (Product, error)
	UpdateProduct(db *sql.DB, product Product) error
	DeleteProduct(db *sql.DB, id int) error
	CreateProduct(db *sql.DB, product Product) (int, error)
	GetProducts(db *sql.DB, start, count int) ([]Product, error)
}
