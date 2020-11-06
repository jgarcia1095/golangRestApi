package customers

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/golangRestApi/helper"
)

type getCustomersRequest struct {
	Limit  int
	Offset int
}

// @Summary Lista de clientes
// @Tags customers
// @Accept json
// @Produce json
// @Param request body customers.getCustomersRequest true "User Data"
// @Success 200 {object} customers.CustomerList "ok"
// @Router /customers/paginated [post]
func makeGetCustomersEndPoint(s Service) endpoint.Endpoint {
	getCustomerEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCustomersRequest)
		result, err := s.GetCustomers(&req)
		helper.Catch(err)
		return result, nil
	}
	return getCustomerEndPoint
}
