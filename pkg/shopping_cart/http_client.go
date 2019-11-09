package shopping_cart

import (
	"fmt"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPClient(instance string) (ShoppingCartService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	addCartEndpoint := httptransport.NewClient(
		"POST",
		copyURL(u, "/add_cart/"),
		encodeHTTPGenericRequest,
		decodeHTTPAddCartResponse,
	).Endpoint()

	addItemEndpoint := httptransport.NewClient(
		"POST",
		copyURL(u, "/add_item/"),
		encodeHTTPGenericRequest,
		decodeHTTPAddItemResponse,
	).Endpoint()

	return Endpoints{
		AddCartEndpoint: addCartEndpoint,
		AddItemEndpoint: addItemEndpoint,
	}, nil
}

func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	fmt.Println(r.Body)
	return nil
}

func decodeHTTPAddCartResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp AddCartResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func decodeHTTPAddItemResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp AddItemResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
