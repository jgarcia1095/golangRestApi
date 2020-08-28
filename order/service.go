package order

import "github.com/golangRestApi/helper"

type Service interface {
	GetOrderByID(params *getOrderByIDRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) (*OrderList, error)
	InsertOrder(params *addOrderRequest) (int64, error)
	UpdateOrder(params *addOrderRequest) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetOrderByID(params *getOrderByIDRequest) (*OrderItem, error) {
	return s.repo.GetOrderByID(params)
}

func (s *service) GetOrders(params *getOrdersRequest) (*OrderList, error) {
	orders, err := s.repo.GetOrders(params)
	helper.Catch(err)

	totalOrders, err := s.repo.GetTotalOrders(params)
	helper.Catch(err)

	return &OrderList{Data: orders, TotalRecords: totalOrders}, nil
}

func (s *service) InsertOrder(params *addOrderRequest) (int64, error) {
	orderID, err := s.repo.InsertOrder(params)
	helper.Catch(err)

	for _, detail := range params.OrderDetails {
		detail.OrderID = orderID
		_, err := s.repo.InsertOrderDetail(&detail)
		helper.Catch(err)
	}

	return orderID, nil
}

func (s *service) UpdateOrder(params *addOrderRequest) (int64, error) {
	orderID, err := s.repo.UpdateOrder(params)
	helper.Catch(err)

	for _, detail := range params.OrderDetails {
		detail.OrderID = orderID
		if detail.ID == 0 {
			_, err = s.repo.InsertOrderDetail(&detail)
		} else {
			_, err = s.repo.UpdateOrderDetail(&detail)
		}
		helper.Catch(err)
	}

	return orderID, nil
}
