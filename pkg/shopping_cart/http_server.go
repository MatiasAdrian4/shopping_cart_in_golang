package shopping_cart

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/examples/addsvc/pkg/addservice"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(endpoints Endpoints) http.Handler {
	
	m := mux.NewRouter()
	
	m.Methods("POST").Path("/add_cart/").Handler(httptransport.NewServer(
		endpoints.AddCartEndpoint,
		decodeHTTPAddCartRequest,
		encodeHTTPGenericResponse,
	))
	m.Methods("GET").Path("/get_cart/{id}").Handler(httptransport.NewServer(
		endpoints.GetCartEndpoint,
		decodeHTTPGetCartRequest,
		encodeHTTPGenericResponse,
	))
	m.Methods("GET").Path("/list_carts/").Handler(httptransport.NewServer(
		endpoints.ListCartsEndpoint,
		decodeHTTPListCartsRequest,
		encodeHTTPGenericResponse,
	))

	m.Methods("POST").Path("/add_item/").Handler(httptransport.NewServer(
		endpoints.AddItemEndpoint,
		decodeHTTPAddItemRequest,
		encodeHTTPGenericResponse,
	))
	m.Methods("GET").Path("/get_item/{id}").Handler(httptransport.NewServer(
		endpoints.GetItemEndpoint,
		decodeHTTPGetItemRequest,
		encodeHTTPGenericResponse,
	))
	m.Methods("GET").Path("/list_items/").Handler(httptransport.NewServer(
		endpoints.ListItemsEndpoint,
		decodeHTTPListItemsRequest,
		encodeHTTPGenericResponse,
	))
	m.Methods("POST").Path("/add_cart_element/").Handler(httptransport.NewServer(
		endpoints.AddCartElementEndpoint,
		decodeHTTPAddCartElementRequest,
		encodeHTTPGenericResponse,
	))

	return m
}


func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func err2code(err error) int {
	switch err {
	case addservice.ErrTwoZeroes, addservice.ErrMaxSizeExceeded, addservice.ErrIntOverflow:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

func decodeHTTPAddCartRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AddCartRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPGetCartRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetCartRequest
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return nil, err
	}
	req.Id = id
	return req, nil
}

func decodeHTTPListCartsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return ListCartsRequest{}, nil
}

func decodeHTTPAddItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AddItemRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPGetItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetItemRequest
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return nil, err
	}
	req.Id = id
	return req, nil
}

func decodeHTTPListItemsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return ListItemsRequest{}, nil
}

func decodeHTTPAddCartElementRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AddCartElementRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
