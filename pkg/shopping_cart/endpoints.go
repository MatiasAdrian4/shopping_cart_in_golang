package shopping_cart

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"shopping_cart_in_golang_with_go_kit/pb"
)

type Endpoints struct {
	AddCartEndpoint 		endpoint.Endpoint
	GetCartEndpoint 		endpoint.Endpoint
	ListCartsEndpoint		endpoint.Endpoint

	AddItemEndpoint 		endpoint.Endpoint
	GetItemEndpoint			endpoint.Endpoint
	ListItemsEndpoint		endpoint.Endpoint

	AddCartElementEndpoint	endpoint.Endpoint
	ListItemsByCartEndpoint	endpoint.Endpoint
}

func MakeEndpoints(svc ShoppingCartService) Endpoints {
	return Endpoints{
		AddCartEndpoint: 			MakeAddCartEndpoint(svc),
		GetCartEndpoint: 			MakeGetCartEndpoint(svc),
		ListCartsEndpoint: 			MakeListCartsEndpoint(svc),

		AddItemEndpoint: 			MakeAddItemEndpoint(svc),
		GetItemEndpoint:			MakeGetItemEndpoint(svc),
		ListItemsEndpoint:			MakeListItemsEndpoint(svc),

		AddCartElementEndpoint: 	MakeAddCartElementEndpoint(svc),
		ListItemsByCartEndpoint:	MakeListItemsByCartEndpoint(svc),
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

func (s Endpoints) ListCarts(ctx context.Context) ([]*pb.Cart, error) {
	resp, err := s.ListCartsEndpoint(ctx, ListCartsRequest{})
	if err != nil {
		return nil, err
	}
	response := resp.(ListCartsResponse)

	return response.Carts, response.Err
}

func (s Endpoints) AddItem(ctx context.Context, Id int, Detail string, Price float64) (int, error) {
	resp, err := s.AddItemEndpoint(ctx, AddItemRequest{Id: Id, Detail: Detail, Price: Price})
	if err != nil {
		return 0, err
	}
	response := resp.(AddItemResponse)

	return response.Id, response.Err
}

func (s Endpoints) GetItem(ctx context.Context, Id int) (*pb.Item, error) {
	resp, err := s.GetCartEndpoint(ctx, GetItemRequest{Id: Id})
	if err != nil {
		return &pb.Item{}, err
	}
	response := resp.(GetItemResponse)

	return response.Item, response.Err
}

func (s Endpoints) ListItems(ctx context.Context) ([]*pb.Item, error) {
	resp, err := s.ListItemsEndpoint(ctx, ListItemsRequest{})
	if err != nil {
		return nil, err
	}
	response := resp.(ListItemsResponse)

	return response.Items, response.Err
}

func (s Endpoints) AddCartElement(ctx context.Context, Cart_id int, Item_id int, Quantity float64) (error) {
	resp, err := s.AddCartElementEndpoint(ctx, AddCartElementRequest{Cart_id: Cart_id, Item_id: Item_id, Quantity: Quantity})
	if err != nil {
		return err
	}
	response := resp.(AddCartElementResponse)

	return response.Err
}

func (s Endpoints) ListItemsByCart(ctx context.Context, Cart_id int) ([]*pb.Item, error) {
	resp, err := s.ListItemsByCartEndpoint(ctx, ListItemsByCartRequest{Cart_id: Cart_id})
	if err != nil {
		return nil, err
	}
	response := resp.(ListItemsByCartResponse)

	return response.Items, response.Err
}

func MakeAddCartEndpoint(svc ShoppingCartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddCartRequest)
		v, err := svc.AddCart(ctx, req.Id)
		return AddCartResponse{v, err}, err
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
		return GetCartResponse{v, err}, err
	}
}

type GetCartRequest struct {
	Id int `json:"id"`
}

type GetCartResponse struct {
	Cart	*pb.Cart	`json:"cart"`
	Err		error 		`json:"err"`
}

func MakeListCartsEndpoint(svc ShoppingCartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.ListCarts(ctx)
		return ListCartsResponse{v, err}, err
	}
}

type ListCartsRequest struct {
}

type ListCartsResponse struct {
	Carts	[]*pb.Cart	`json:"carts"`
	Err		error 		`json:"err"`
}

func MakeAddItemEndpoint(svc ShoppingCartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddItemRequest)
		v, err := svc.AddItem(ctx, req.Id, req.Detail, req.Price)
		return AddItemResponse{v, err}, err
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

func MakeGetItemEndpoint(svc ShoppingCartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetItemRequest)
		v, err := svc.GetItem(ctx, req.Id)
		return GetItemResponse{v, err}, err
	}
}

type GetItemRequest struct {
	Id int `json:"id"`
}

type GetItemResponse struct {
	Item	*pb.Item	`json:"item"`
	Err		error 		`json:"err"`
}

func MakeListItemsEndpoint(svc ShoppingCartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.ListItems(ctx)
		return ListItemsResponse{v, err}, err
	}
}

type ListItemsRequest struct {
}

type ListItemsResponse struct {
	Items	[]*pb.Item	`json:"items"`
	Err		error 		`json:"err"`
}

func MakeAddCartElementEndpoint(svc ShoppingCartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddCartElementRequest)
		err := svc.AddCartElement(ctx, req.Cart_id, req.Item_id, req.Quantity)
		return AddCartElementResponse{err}, err
	}
}

type AddCartElementRequest struct {
	Cart_id     int     `json:"cart_id"`
	Item_id 	int  	`json:"item_id"`
	Quantity  	float64 `json:"quantity"`
}

type AddCartElementResponse struct {
	Err 	error `json:"err"`
}

func MakeListItemsByCartEndpoint(svc ShoppingCartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListItemsByCartRequest)
		v, err := svc.ListItemsByCart(ctx, req.Cart_id)
		if err != nil {
			return ListItemsByCartResponse{v, err}, nil
		}
		return ListItemsByCartResponse{v, nil}, nil
	}
}

type ListItemsByCartRequest struct {
	Cart_id	int	`json:"cart_id"`
}

type ListItemsByCartResponse struct {
	Items	[]*pb.Item	`json:"items"`
	Err		error 		`json:"err"`
}