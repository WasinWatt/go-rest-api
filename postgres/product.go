package postgres

import (
	"database/sql"

	"github.com/wasinwatt/go-postgres-api/product"
	// "errors"
)

func NewProductRepository() *productRepository {
	return &productRepository{}
}

type productRepository struct{}

func (productRepository) getProduct(db *sql.DB) (product.Product, error) {
	var p product.Product
	err := db.QueryRow("SELECT name, price FROM products WHERE id=$1", p.ID).Scan(&p)

	return p, err
}

func (productRepository) updateProduct(db *sql.DB) error {
	var p product.Product
	_, err := db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3", p.Name, p.Price, p.ID)

	return err
}

func (productRepository) deleteProduct(db *sql.DB) error {
	var p product.Product
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

	return err
}

func (productRepository) createProduct(db *sql.DB, P product.Product) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO products(name, price) VALUES($1, $2) RETURNING id", P.Name, P.Price).Scan(id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (productRepository) getProducts(db *sql.DB, start, count int) ([]product.Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products LIMIT $1 OFFSET $2", count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product.Product{}

	for rows.Next() {
		var p product.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
