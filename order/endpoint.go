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

type addOrderRequest struct {
	ID           int64
	OrderDate    string
	CustomerID   int
	OrderDetails []addOrderDetailRequest
}

type addOrderDetailRequest struct {
	ID        int64
	OrderID   int64
	ProductID int64
	Quantity  int64
	UnitPrice float64
}

type deleteOrderDetailRequest struct {
	OrderDetailID string
}

type deleteOrderRequest struct {
	OrderID string
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

func makeInsertOrderEndPoint(s Service) endpoint.Endpoint {
	insertOrderEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		result, err := s.InsertOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return insertOrderEndPoint
}

func makeUpdateOrderEndPoint(s Service) endpoint.Endpoint {
	updateOrderEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		result, err := s.UpdateOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return updateOrderEndPoint
}

func makeDeleteOrderDetailEndPoint(s Service) endpoint.Endpoint {
	deleteOrderDetailEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderDetailRequest)
		r, err := s.DeleteOrderDetail(&req)
		helper.Catch(err)
		return r, nil
	}
	return deleteOrderDetailEndPoint
}

func makeDeleteOrderEndPoint(s Service) endpoint.Endpoint {
	deleteOrderEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderRequest)
		r, err := s.DeleteOrder(&req)
		helper.Catch(err)
		return r, nil
	}
	return deleteOrderEndPoint
}
