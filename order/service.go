package order

import "github.com/golangRestApi/helper"

type Service interface {
	GetOrderByID(params *getOrderByIDRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) (*OrderList, error)
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
