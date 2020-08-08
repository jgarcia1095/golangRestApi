package main

import (
	"net/http"

	"github.com/golangRestApi/database"
	"github.com/golangRestApi/employee"
	"github.com/golangRestApi/product"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataBaseConnection := database.InitDb()
	defer dataBaseConnection.Close()

	var (
		productRepository  = product.NewRepository(dataBaseConnection)
		employeeRepository = employee.NewRepository(dataBaseConnection)
	)

	var (
		productService  product.Service
		employeeService employee.Service
	)

	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHTTPHandler(productService))
	r.Mount("/employee", employee.MakeHTTPHandler(employeeService))
	http.ListenAndServe(":3000", r)
}
