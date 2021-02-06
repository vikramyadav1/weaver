package main

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/vikramyadav1/weaver/sampleApp/api"
	"github.com/vikramyadav1/weaver/sampleApp/models/product"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("cannot connect to the database")
	}
	defer db.Close()

	pr := product.SqlProductRepository{Db: db}
	ps := api.ProductServer{Pr: &pr}
	ps.RegisterRoutes(r)

	http.Handle("/", r)

	http.ListenAndServe(":3000", nil)
}
