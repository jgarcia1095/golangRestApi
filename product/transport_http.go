package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()
	getProductByIdHandler := kithttp.NewServer(makeGetProductByIdEndPoint(s), getProductByIdRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductByIdHandler)

	getProductsHandler := kithttp.NewServer(makeGetProductsEndPoint(s), getProductsRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	addProductHandler := kithttp.NewServer(makeAddProductsEndPoint(s), addProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addProductHandler)

	getBestEmployeeIdHandler := kithttp.NewServer(makeBestEmmployeeEndPoint(s), getBestEmployeeRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/best", getBestEmployeeIdHandler)
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

func getBestEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getBestEmployeeRequest{}, nil
}
