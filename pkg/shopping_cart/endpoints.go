package shopping_cart

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddCartEndpoint endpoint.Endpoint
	AddItemEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc CartService) Endpoints {
	var addCartEndpoint endpoint.Endpoint
	addCartEndpoint = MakeAddCartEndpoint(svc)

	var addItemEndpoint endpoint.Endpoint
	addItemEndpoint = MakeAddItemEndpoint(svc)

	return Endpoints{
		AddCartEndpoint: addCartEndpoint,
		AddItemEndpoint: addItemEndpoint,
	}
}

func (s Endpoints) AddCart(ctx context.Context, Id int) (int, error) {
	resp, err := s.AddCartEndpoint(ctx, AddCartRequest{Id: Id})
	if err != nil {
		return 0, err
	}
	response := resp.(AddCartResponse)

	return response.Id, response.Err
}

func (s Endpoints) AddItem(ctx context.Context, Id int, Detail string, Price float64) (int, error) {
	resp, err := s.AddItemEndpoint(ctx, AddItemRequest{Id: Id, Detail: Detail, Price: Price})
	if err != nil {
		return 0, err
	}
	response := resp.(AddItemResponse)

	return response.Id, response.Err
}

func MakeAddCartEndpoint(svc CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddCartRequest)
		v, err := svc.AddCart(ctx, req.Id)
		if err != nil {
			return AddCartResponse{v, err}, nil
		}
		return AddCartResponse{v, nil}, nil
	}
}

func MakeAddItemEndpoint(svc CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddItemRequest)
		v, err := svc.AddItem(ctx, req.Id, req.Detail, req.Price)
		if err != nil {
			return AddItemResponse{v, err}, nil
		}
		return AddItemResponse{v, nil}, nil
	}
}

type AddCartRequest struct {
	Id int `json:"id"`
}

type AddCartResponse struct {
	Id  int   `json:"id"`
	Err error `json:"err"`
}

type AddItemRequest struct {
	Id     int     `json:"id"`
	Detail string  `json:"detail"`
	Price  float64 `json:"price"`
}

type AddItemResponse struct {
	Id  int   `json:"id"`
	Err error `json:"err"`
}
