package postgres

import (
	"database/sql"

	"github.com/wasinwatt/go-rest-api/product"
)

func NewProductRepository() product.Repository {
	return &productRepository{}
}

type productRepository struct{}

type productView struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Price string `db:"price"`
}

func (productRepository) GetProduct(db *sql.DB, id int) (product.Product, error) {
	var p product.Product
	err := db.QueryRow("SELECT name, price FROM products WHERE id=$1", id).Scan(&p)

	return p, err
}

func (productRepository) UpdateProduct(db *sql.DB, P product.Product) error {
	_, err := db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3", P.Name, P.Price, P.ID)

	return err
}

func (productRepository) DeleteProduct(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", id)

	return err
}

func (productRepository) CreateProduct(db *sql.DB, P product.Product) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO products(name, price) VALUES($1, $2) RETURNING id", P.Name, P.Price).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (productRepository) GetProducts(db *sql.DB, start, count int) ([]product.Product, error) {
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
