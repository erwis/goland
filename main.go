package main

import (
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goland/database"
	"github.com/goland/product"
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
