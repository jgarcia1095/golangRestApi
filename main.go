package main

import (
	"net/http"

	"github.com/golangRestApi/customers"
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
		productRepository   = product.NewRepository(dataBaseConnection)
		employeeRepository  = employee.NewRepository(dataBaseConnection)
		customersRepository = customers.NewRepository(dataBaseConnection)
	)

	var (
		productService   product.Service
		employeeService  employee.Service
		customersService customers.Service
	)

	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)
	customersService = customers.NewService(customersRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHTTPHandler(productService))
	r.Mount("/employee", employee.MakeHTTPHandler(employeeService))
	r.Mount("/customers", customers.MakeHTTPHandler(customersService))
	http.ListenAndServe(":3000", r)
}
