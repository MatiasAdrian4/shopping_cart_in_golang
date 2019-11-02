package shopping_cart

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/examples/addsvc/pkg/addservice"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(endpoints Endpoints) http.Handler {

	m := http.NewServeMux()
	m.Handle("/add_cart", httptransport.NewServer(
		endpoints.AddCartEndpoint,
		decodeHTTPAddCartRequest,
		encodeHTTPGenericResponse,
	))
	m.Handle("/add_item", httptransport.NewServer(
		endpoints.AddItemEndpoint,
		decodeHTTPAddItemRequest,
		encodeHTTPGenericResponse,
	))

	return m
}

func NewHTTPClient(instance string) (CartService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	var addCartEndpoint endpoint.Endpoint
	addCartEndpoint = httptransport.NewClient(
		"POST",
		copyURL(u, "/add_cart"),
		encodeHTTPGenericRequest,
		decodeHTTPAddCartResponse,
	).Endpoint()

	var addItemEndpoint endpoint.Endpoint
	addItemEndpoint = httptransport.NewClient(
		"POST",
		copyURL(u, "/add_cart"),
		encodeHTTPGenericRequest,
		decodeHTTPAddItemResponse,
	).Endpoint()

	return Endpoints{
		AddCartEndpoint: addCartEndpoint,
		AddItemEndpoint: addItemEndpoint,
	}, nil
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

func decodeHTTPAddCartResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp AddCartResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func decodeHTTPAddItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AddItemRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeHTTPAddItemResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp AddItemResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
