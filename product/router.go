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
	router.HandleFunc("/", handler.getProducts).Methods("GET")
	router.HandleFunc("/create", handler.createProduct).Methods("POST")
	return &ProductRouter{router, db}
}