package employee

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/golangRestApi/helper"
)

func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getEmployeeByIDHandler := kithttp.NewServer(makeGetEmployeeByIDEndPoint(s), getEmployeeByIDRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getEmployeeByIDHandler)

	getEmployeesHandler := kithttp.NewServer(makeGetEmployeesEndPoint(s), getEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getEmployeesHandler)

	getBestEmployeeIDHandler := kithttp.NewServer(makeBestEmployeeEndPoint(s), getBestEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/best", getBestEmployeeIDHandler)

	addEmployeeIDHandler := kithttp.NewServer(makeAddEmployeeEndPoint(s), addEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addEmployeeIDHandler)

	updateEmployeeHandler := kithttp.NewServer(makeUpdateEmployeeEndPoint(s), updateEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateEmployeeHandler)

	deleteEmployeeHandler := kithttp.NewServer(makeDeleteEmployeeEndPoint(s), deleteEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteEmployeeHandler)

	return r
}

func getEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getEmployeeByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getEmployeeByIDRequest{
		EmployeeID: chi.URLParam(r, "id"),
	}, nil
}

func getBestEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getBestEmployeeRequest{}, nil
}

func addEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addEmployeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func updateEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateEmployeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func deleteEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return deleteEmployeeRequest{
		EmployeeID: chi.URLParam(r, "id"),
	}, nil
}
