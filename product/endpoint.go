package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIdRequest struct {
	ProductId int
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

type getBestEmployeeRequest struct {
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIdRequest)
		product, err := s.GetProductById(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}
	return getProductByIdEndPoint
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

func makeBestEmmployeeEndPoint(s Service) endpoint.Endpoint {
	getBestEmployeeEndPoint := func(_ context.Context, _ interface{}) (interface{}, error) {
		result, err := s.GetBestsEmployee()
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getBestEmployeeEndPoint
}
