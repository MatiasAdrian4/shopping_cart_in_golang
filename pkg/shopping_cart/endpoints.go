package shopping_cart

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"shopping_cart/pb"
)

type Endpoints struct {
	AddCartEndpoint endpoint.Endpoint
	GetCartEndpoint endpoint.Endpoint
	AddItemEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc ShoppingCartService) Endpoints {
	return Endpoints{
		AddCartEndpoint: MakeAddCartEndpoint(svc),
		GetCartEndpoint: MakeGetCartEndpoint(svc),
		AddItemEndpoint: MakeAddItemEndpoint(svc),
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

func (s Endpoints) GetCart(ctx context.Context, Id int) (*pb.Cart, error) {
	resp, err := s.GetCartEndpoint(ctx, GetCartRequest{Id: Id})
	if err != nil {
		return &pb.Cart{}, err
	}
	response := resp.(GetCartResponse)

	return response.Cart, response.Err
}

func (s Endpoints) AddItem(ctx context.Context, Id int, Detail string, Price float64) (int, error) {
	resp, err := s.AddItemEndpoint(ctx, AddItemRequest{Id: Id, Detail: Detail, Price: Price})
	if err != nil {
		return 0, err
	}
	response := resp.(AddItemResponse)

	return response.Id, response.Err
}

func MakeAddCartEndpoint(svc ShoppingCartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddCartRequest)
		v, err := svc.AddCart(ctx, req.Id)
		if err != nil {
			return AddCartResponse{v, err}, nil
		}
		return AddCartResponse{v, nil}, nil
	}
}

type AddCartRequest struct {
	Id int `json:"id"`
}

type AddCartResponse struct {
	Id  int   `json:"id"`
	Err error `json:"err"`
}

func MakeGetCartEndpoint(svc ShoppingCartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCartRequest)
		v, err := svc.GetCart(ctx, req.Id)
		if err != nil {
			return GetCartResponse{v, err}, nil
		}
		return GetCartResponse{v, nil}, nil
	}
}

type GetCartRequest struct {
	Id int `json:"id"`
}

type GetCartResponse struct {
	Cart  *pb.Cart   `json:"cart"`
	Err	error `json:"err"`
}

func MakeAddItemEndpoint(svc ShoppingCartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddItemRequest)
		v, err := svc.AddItem(ctx, req.Id, req.Detail, req.Price)
		if err != nil {
			return AddItemResponse{v, err}, nil
		}
		return AddItemResponse{v, nil}, nil
	}
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
