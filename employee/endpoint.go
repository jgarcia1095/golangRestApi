package employee

import (
	"context"

	"github.com/golangRestApi/helper"

	"github.com/go-kit/kit/endpoint"
)

type getEmployeeByIDRequest struct {
	EmployeeID string
}

type getEmployeeRequest struct {
	Limit  int
	Offset int
}

// type getAddProductsRequest struct {
// 	Category     string
// 	Description  string
// 	ListPrice    string
// 	StandardCost string
// 	ProductCode  string
// 	ProductName  string
// }

// type updateProductsRequest struct {
// 	ID           int64
// 	Category     string
// 	Description  string
// 	ListPrice    float32
// 	StandardCost float32
// 	ProductCode  string
// 	ProductName  string
// }

// type deleteProductsRequest struct {
// 	ProductID string
// }

// type getBestSellersRequest struct {
// }

// type getBestEmployeeRequest struct {
// }

// type addEmployeesRequest struct {
// 	Address       string
// 	BusinessPhone string
// 	Company       string
// 	EmailAddress  string
// 	FaxNumber     string
// 	FirstName     string
// 	HomePhone     string
// 	JobTitle      string
// 	LasttName     string
// 	MobilPhone    string
// }

func makeGetEmployeeByIDEndPoint(s Service) endpoint.Endpoint {
	getEmployeeByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeeByIDRequest)
		employee, err := s.GetEmployeeByID(&req)
		helper.Catch(err)
		return employee, nil
	}
	return getEmployeeByIDEndPoint
}

func makeGetEmployeesEndPoint(s Service) endpoint.Endpoint {
	getEmployeesEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeeRequest)
		result, err := s.GetEmployees(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getEmployeesEndPoint
}

// func makeAddProductsEndPoint(s Service) endpoint.Endpoint {
// 	addProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(getAddProductsRequest)
// 		productID, err := s.InsertProduct(&req)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return productID, nil
// 	}
// 	return addProductsEndPoint
// }

// // func makeBestEmmployeeEndPoint(s Service) endpoint.Endpoint {
// // 	getBestEmployeeEndPoint := func(_ context.Context, _ interface{}) (interface{}, error) {
// // 		result, err := s.GetBestsEmployee()
// // 		if err != nil {
// // 			panic(err)
// // 		}
// // 		return result, nil
// // 	}
// // 	return getBestEmployeeEndPoint
// // }

// func makeAddEmployeeEndPoint(s Service) endpoint.Endpoint {
// 	addEmployeeEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(addEmployeesRequest)
// 		employeeID, err := s.InsertEmployee(&req)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return employeeID, nil
// 	}
// 	return addEmployeeEndPoint
// }

// func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {
// 	updateProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(updateProductsRequest)
// 		r, err := s.UpdateProduct(&req)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return r, nil
// 	}
// 	return updateProductsEndPoint
// }

// func makeDeleteProductEndPoint(s Service) endpoint.Endpoint {
// 	deleteProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(deleteProductsRequest)
// 		r, err := s.DeleteProduct(&req)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return r, nil
// 	}
// 	return deleteProductsEndPoint
// }

// func makeBestSellersEndPoint(s Service) endpoint.Endpoint {
// 	getBestSellersEndPoint := func(_ context.Context, _ interface{}) (interface{}, error) {
// 		result, err := s.GetBestSellers()
// 		if err != nil {
// 			panic(err)
// 		}
// 		return result, nil
// 	}
// 	return getBestSellersEndPoint
// }
