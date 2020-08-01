package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()
	getProductByIDHandler := kithttp.NewServer(makeGetProductByIdEndPoint(s), getProductByIdRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductByIDHandler)

	getProductsHandler := kithttp.NewServer(makeGetProductsEndPoint(s), getProductsRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	addProductHandler := kithttp.NewServer(makeAddProductsEndPoint(s), addProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addProductHandler)

	updateProductHandler := kithttp.NewServer(makeUpdateProductEndPoint(s), updateProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateProductHandler)

	deleteProductHandler := kithttp.NewServer(makeDeleteProductEndPoint(s), deleteProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteProductHandler)

	getBestEmployeeIDHandler := kithttp.NewServer(makeBestEmmployeeEndPoint(s), getBestEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/best", getBestEmployeeIDHandler)

	// addEmployeeIdHandler := kithttp.NewServer(makeAddEmployeeEndPoint(s), addEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	// r.Method(http.MethodPost, "/", addEmployeeIdHandler)

	return r
}

func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func getProductByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIdRequest{
		ProductId: productId,
	}, nil
}

func addProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func updateProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}

func deleteProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return deleteProductsRequest{
		ProductID: chi.URLParam(r, "id"),
	}, nil
}

func getBestEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getBestEmployeeRequest{}, nil
}

func addEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addEmployeesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}
