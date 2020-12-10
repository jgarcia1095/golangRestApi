package main

import (
	"net/http"

	"github.com/golangRestApi/customers"
	"github.com/golangRestApi/database"
	"github.com/golangRestApi/employee"
	"github.com/golangRestApi/helper"
	"github.com/golangRestApi/order"
	"github.com/golangRestApi/product"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golangRestApi/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Northwind API
// @version 1.0
// @description This is a sampe server celler server
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func main() {
	dataBaseConnection := database.InitDb()
	defer dataBaseConnection.Close()

	var (
		productRepository   = product.NewRepository(dataBaseConnection)
		employeeRepository  = employee.NewRepository(dataBaseConnection)
		customersRepository = customers.NewRepository(dataBaseConnection)
		orderRepository     = order.NewRepository(dataBaseConnection)
	)

	var (
		productService   product.Service
		employeeService  employee.Service
		customersService customers.Service
		orderService     order.Service
	)

	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)
	customersService = customers.NewService(customersRepository)
	orderService = order.NewService(orderRepository)

	r := chi.NewRouter()
	r.Use(helper.GetCors().Handler)
	r.Mount("/products", product.MakeHTTPHandler(productService))
	r.Mount("/employee", employee.MakeHTTPHandler(employeeService))
	r.Mount("/customers", customers.MakeHTTPHandler(customersService))
	r.Mount("/orders", order.MakeHTTPHandler(orderService))
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("../swagger/doc.json"),
	))
	http.ListenAndServe(":3000", r)
}
