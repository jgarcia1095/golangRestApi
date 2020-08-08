package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/golangRestApi/helper"
)

func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()
	getProductByIDHandler := kithttp.NewServer(makeGetProductByIDEndPoint(s), getProductByIDRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductByIDHandler)

	getProductsHandler := kithttp.NewServer(makeGetProductsEndPoint(s), getProductsRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	addProductHandler := kithttp.NewServer(makeAddProductsEndPoint(s), addProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addProductHandler)

	updateProductHandler := kithttp.NewServer(makeUpdateProductEndPoint(s), updateProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateProductHandler)

	deleteProductHandler := kithttp.NewServer(makeDeleteProductEndPoint(s), deleteProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteProductHandler)

	getBestSellerHandler := kithttp.NewServer(makeBestSellersEndPoint(s), getBestSellersRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/bestSellers", getBestSellerHandler)

	// getBestEmployeeIDHandler := kithttp.NewServer(makeBestEmmployeeEndPoint(s), getBestEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	// r.Method(http.MethodGet, "/bestSellers", getBestEmployeeIDHandler)

	// addEmployeeIdHandler := kithttp.NewServer(makeAddEmployeeEndPoint(s), addEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	// r.Method(http.MethodPost, "/", addEmployeeIdHandler)

	return r
}

func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getProductByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIDRequest{
		ProductID: productID,
	}, nil
}

func addProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func updateProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
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

func getBestSellersRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getBestSellersRequest{}, nil
}

func addEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addEmployeesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}
