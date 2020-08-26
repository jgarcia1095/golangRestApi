package order

import (
	"context"

	"github.com/golangRestApi/helper"

	"github.com/go-kit/kit/endpoint"
)

type getOrderByIDRequest struct {
	OrderID int64
}

type getOrdersRequest struct {
	Limit    int
	Offset   int
	Status   interface{}
	DateFrom interface{}
	DateTo   interface{}
}

func makeGetOrdersByIDEndPoint(s Service) endpoint.Endpoint {
	getOrdersByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrderByIDRequest)
		result, err := s.GetOrderByID(&req)
		helper.Catch(err)
		return result, nil
	}
	return getOrdersByIDEndPoint
}

func makeGetOrdersEndPoint(s Service) endpoint.Endpoint {
	getOrdersEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrdersRequest)
		result, err := s.GetOrders(&req)
		helper.Catch(err)
		return result, nil
	}
	return getOrdersEndPoint
}
