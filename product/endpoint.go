package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getAddProductsRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

type updateProductsRequest struct {
	ID           int64
	Category     string
	Description  string
	ListPrice    float32
	StandardCost float32
	ProductCode  string
	ProductName  string
}

type deleteProductsRequest struct {
	ProductID string
}

type getBestSellersRequest struct {
}

type getBestEmployeeRequest struct {
}

type addEmployeesRequest struct {
	Address       string
	BusinessPhone string
	Company       string
	EmailAddress  string
	FaxNumber     string
	FirstName     string
	HomePhone     string
	JobTitle      string
	LasttName     string
	MobilPhone    string
}

func makeGetProductByIDEndPoint(s Service) endpoint.Endpoint {
	getProductByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductByID(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}
	return getProductByIDEndPoint
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getProductsEndPoint
}

func makeAddProductsEndPoint(s Service) endpoint.Endpoint {
	addProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductsRequest)
		productID, err := s.InsertProduct(&req)
		if err != nil {
			panic(err)
		}
		return productID, nil
	}
	return addProductsEndPoint
}

func makeAddEmployeeEndPoint(s Service) endpoint.Endpoint {
	addEmployeeEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addEmployeesRequest)
		employeeID, err := s.InsertEmployee(&req)
		if err != nil {
			panic(err)
		}
		return employeeID, nil
	}
	return addEmployeeEndPoint
}

func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {
	updateProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductsRequest)
		r, err := s.UpdateProduct(&req)
		if err != nil {
			panic(err)
		}
		return r, nil
	}
	return updateProductsEndPoint
}

func makeDeleteProductEndPoint(s Service) endpoint.Endpoint {
	deleteProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductsRequest)
		r, err := s.DeleteProduct(&req)
		if err != nil {
			panic(err)
		}
		return r, nil
	}
	return deleteProductsEndPoint
}

func makeBestSellersEndPoint(s Service) endpoint.Endpoint {
	getBestSellersEndPoint := func(_ context.Context, _ interface{}) (interface{}, error) {
		result, err := s.GetBestSellers()
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getBestSellersEndPoint
}
