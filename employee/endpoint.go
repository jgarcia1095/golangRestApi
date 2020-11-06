package employee

import (
	"context"

	"github.com/golangRestApi/helper"

	"github.com/go-kit/kit/endpoint"
)

type getEmployeeByIDRequest struct {
	EmployeeID string
}

type getEmployeeRequest struct {
	Limit  int
	Offset int
}

type getBestEmployeeRequest struct {
}

type addEmployeeRequest struct {
	FirstName     string
	LastName      string
	Address       string
	EmailAddress  string
	MobilePhone   string
	HomePhone     string
	BusinessPhone string
	FaxNumber     string
	JobTitle      string
	Company       string
}

type updateEmployeeRequest struct {
	ID            int64
	FirstName     string
	LastName      string
	Address       string
	EmailAddress  string
	MobilePhone   string
	HomePhone     string
	BusinessPhone string
	FaxNumber     string
	JobTitle      string
	Company       string
}

type deleteEmployeeRequest struct {
	EmployeeID string
}

// @Summary Emmpleado po Id
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path int true "Employee Id"
// @Success 200 {object} employee.Employee "ok"
// @Router /employee/{id} [get]
func makeGetEmployeeByIDEndPoint(s Service) endpoint.Endpoint {
	getEmployeeByIDEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeeByIDRequest)
		employee, err := s.GetEmployeeByID(&req)
		helper.Catch(err)
		return employee, nil
	}
	return getEmployeeByIDEndPoint
}

// @Summary Lista de Empleados
// @Tags Employee
// @Accept json
// @Produce json
// @Param request body employee.getEmployeeRequest true "User Data"
// @Success 200 {object} employee.EmployeeList "ok"
// @Router /employee/paginated [post]
func makeGetEmployeesEndPoint(s Service) endpoint.Endpoint {
	getEmployeesEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeeRequest)
		result, err := s.GetEmployees(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getEmployeesEndPoint
}

func makeBestEmployeeEndPoint(s Service) endpoint.Endpoint {
	getBestEmployeeEndPoint := func(_ context.Context, _ interface{}) (interface{}, error) {
		result, err := s.GetBestEmployee()
		helper.Catch(err)
		return result, nil
	}
	return getBestEmployeeEndPoint
}

func makeAddEmployeeEndPoint(s Service) endpoint.Endpoint {
	addEmployeeEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addEmployeeRequest)
		employeeID, err := s.InsertEmployee(&req)
		if err != nil {
			panic(err)
		}
		return employeeID, nil
	}
	return addEmployeeEndPoint
}

func makeUpdateEmployeeEndPoint(s Service) endpoint.Endpoint {
	updateEmployeeEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateEmployeeRequest)
		r, err := s.UpdateEmployee(&req)
		helper.Catch(err)
		return r, nil
	}
	return updateEmployeeEndPoint
}

func makeDeleteEmployeeEndPoint(s Service) endpoint.Endpoint {
	deleteEmployeeEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteEmployeeRequest)
		r, err := s.DeleteEmployee(&req)
		helper.Catch(err)
		return r, nil
	}
	return deleteEmployeeEndPoint
}
