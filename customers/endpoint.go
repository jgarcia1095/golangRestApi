package customers

import (
	"context"

	"github.com/golangRestApi/helper"

	"github.com/go-kit/kit/endpoint"
)

type getCustomersRequest struct {
	Limit  int
	Offset int
}

func makeGetCustomersEndPoint(s Service) endpoint.Endpoint {
	getCustomerEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCustomersRequest)
		result, err := s.GetCustomers(&req)
		helper.Catch(err)
		return result, nil
	}
	return getCustomerEndPoint
}
