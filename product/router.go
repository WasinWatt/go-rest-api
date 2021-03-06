package product

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type ProductRouter struct {
	Router *mux.Router
	DB     *sql.DB
}

type ProductHandler struct {
	DB      *sql.DB
	product Repository
}

func NewProductRouter(db *sql.DB, product Repository) *ProductRouter {
	router := mux.NewRouter()
	handler := &ProductHandler{db, product}
	router.HandleFunc("/product", handler.getProducts).Methods("GET")
	router.HandleFunc("/product", handler.createProduct).Methods("POST")
	router.HandleFunc("/product/{id:[0-9]+}", handler.getProduct).Methods("GET")
	router.HandleFunc("/product/{id:[0-9]+}", handler.updateProduct).Methods("PUT")
	router.HandleFunc("/product/{id:[0-9]+}", handler.deleteProduct).Methods("DELETE")

	return &ProductRouter{router, db}
}
