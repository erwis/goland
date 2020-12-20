package main

import (
	"net/http"

	"goland/database"
	"goland/product"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := database.InitDB()
	defer db.Close()
	var productRepository = product.NewRepository(db)
	var productService product.Service
	productService = product.NewService(productRepository)

	r := chi.NewRouter()
	r.Mount("/product", product.MakeHttpHeadler(productService))
	http.ListenAndServe(":3000", r)
}
