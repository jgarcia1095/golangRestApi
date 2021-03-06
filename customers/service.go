package customers

import "github.com/golangRestApi/helper"

type Service interface {
	GetCustomers(params *getCustomersRequest) (*CustomerList, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetCustomers(params *getCustomersRequest) (*CustomerList, error) {
	customers, err := s.repo.GetCustomers(params)

	helper.Catch(err)
	totalCustomers, err := s.repo.GetTotalCustomers()

	helper.Catch(err)
	return &CustomerList{Data: customers, TotalRecords: totalCustomers}, nil
}
