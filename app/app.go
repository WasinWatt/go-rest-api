package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/wasinwatt/go-rest-api/postgres"
	"github.com/wasinwatt/go-rest-api/product"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(url string) {
	var err error
	a.DB, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Println("YEY we're in")
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) initializeRoutes() {
	productRepo := postgres.NewProductRepository()
	productRouter := product.NewProductRouter(a.DB, productRepo)
	a.Router.PathPrefix("/product").Handler(productRouter.Router)
}
