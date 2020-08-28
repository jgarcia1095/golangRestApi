package order

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golangRestApi/helper"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getOrdertByIDHandler := kithttp.NewServer(makeGetOrdersByIDEndPoint(s), getOrderByIDRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getOrdertByIDHandler)

	getOrdersHandler := kithttp.NewServer(makeGetOrdersEndPoint(s), getOrdersRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getOrdersHandler)

	addOrderHandler := kithttp.NewServer(makeInsertOrderEndPoint(s), addOrderRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addOrderHandler)
	return r
}

func getOrderByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	orderID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	helper.Catch(err)

	return getOrderByIDRequest{
		OrderID: orderID,
	}, nil
}

func getOrdersRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getOrdersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func addOrderRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}
