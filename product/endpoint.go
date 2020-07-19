package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIdRequest struct {
	ProductId int
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
