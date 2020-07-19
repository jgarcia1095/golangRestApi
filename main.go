package main

import (
	"net/http"

	"github.com/golangRestApi/database"
	"github.com/golangRestApi/product"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataBaseConnection := database.InitDb()
	defer dataBaseConnection.Close()

	var productRepository = product.NewRepository(dataBaseConnection)
	var productService product.Service
	productService = product.NewService(productRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))
	http.ListenAndServe(":3000", r)
}
