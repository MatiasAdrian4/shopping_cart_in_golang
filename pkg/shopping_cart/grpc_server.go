package shopping_cart

import (
	"context"
	"errors"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	"shopping_cart/pb"
)

type grpcServer struct {
	addCart 		grpctransport.Handler
	getCart 		grpctransport.Handler
	listCarts 		grpctransport.Handler
	addItem 		grpctransport.Handler
	getItem 		grpctransport.Handler
	listItems 		grpctransport.Handler
	addCartElement 	grpctransport.Handler
}

func NewGRPCServer(endpoints Endpoints) pb.ShoppingCartServer {
	return &grpcServer{
		addCart: grpctransport.NewServer(
			endpoints.AddCartEndpoint,
			decodeGRPCAddCartRequest,
			encodeGRPCAddCartResponse,
		),
		getCart: grpctransport.NewServer(
			endpoints.GetCartEndpoint,
			decodeGRPCGetCartRequest,
			encodeGRPCGetCartResponse,
		),
		listCarts: grpctransport.NewServer(
			endpoints.ListCartsEndpoint,
			decodeGRPCListCartsRequest,
			encodeGRPCListCartsResponse,
		),
		addItem: grpctransport.NewServer(
			endpoints.AddItemEndpoint,
			decodeGRPCAddItemRequest,
			encodeGRPCAddItemResponse,
		),
		getItem: grpctransport.NewServer(
			endpoints.GetItemEndpoint,
			decodeGRPCGetItemRequest,
			encodeGRPCGetItemResponse,
		),
		listItems: grpctransport.NewServer(
			endpoints.ListItemsEndpoint,
			decodeGRPCListItemsRequest,
			encodeGRPCListItemsResponse,
		),
		addCartElement: grpctransport.NewServer(
			endpoints.AddCartElementEndpoint,
			decodeGRPCAddCartElementRequest,
			encodeGRPCAddCartElementResponse,
		),
	}
}

func (s *grpcServer) AddCart(ctx context.Context, req *pb.AddCartRequest) (*pb.AddCartResponse, error) {
	_, resp, err := s.addCart.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AddCartResponse), nil
}

func decodeGRPCAddCartRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddCartRequest)
	return AddCartRequest{
		Id: int(req.Id),
	}, nil
}

func encodeGRPCAddCartResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(AddCartResponse)
	return &pb.AddCartResponse{
		Id: int64(resp.Id), 
		Err: err2str(resp.Err),
	}, nil
}

func (s *grpcServer) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	_, resp, err := s.getCart.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetCartResponse), nil
}

func decodeGRPCGetCartRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetCartRequest)
	return GetCartRequest{
		Id: int(req.Id),
	}, nil
}

func encodeGRPCGetCartResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(GetCartResponse)
	return &pb.GetCartResponse{
		Cart: resp.Cart, 
		Err: err2str(resp.Err),
	}, nil
}

func (s *grpcServer) ListCarts(ctx context.Context, req *pb.ListCartsRequest) (*pb.ListCartsResponse, error) {
	_, resp, err := s.listCarts.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ListCartsResponse), nil
}

func decodeGRPCListCartsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return ListCartsRequest{}, nil
}

func encodeGRPCListCartsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(ListCartsResponse)
	return &pb.ListCartsResponse{
		Carts: resp.Carts, 
		Err: err2str(resp.Err),
	}, nil
}

func (s *grpcServer) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	_, resp, err := s.addItem.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AddItemResponse), nil
}

func decodeGRPCAddItemRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddItemRequest)
	return AddItemRequest{
		Id: int(req.Id), 
		Detail: req.Detail, 
		Price: float64(req.Price),
	}, nil
}

func encodeGRPCAddItemResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(AddItemResponse)
	return &pb.AddItemResponse{
		Id: int64(resp.Id), 
		Err: err2str(resp.Err),
	}, nil
}

func (s *grpcServer) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	_, resp, err := s.getItem.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetItemResponse), nil
}

func decodeGRPCGetItemRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetItemRequest)
	return GetItemRequest{
		Id: int(req.Id),
	}, nil
}

func encodeGRPCGetItemResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(GetItemResponse)
	return &pb.GetItemResponse{
		Item: resp.Item, 
		Err: err2str(resp.Err),
	}, nil
}

func (s *grpcServer) ListItems(ctx context.Context, req *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {
	_, resp, err := s.listItems.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ListItemsResponse), nil
}

func decodeGRPCListItemsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return ListItemsRequest{}, nil
}

func encodeGRPCListItemsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(ListItemsResponse)
	return &pb.ListItemsResponse{
		Items: resp.Items, 
		Err: err2str(resp.Err),
	}, nil
}

func (s *grpcServer) AddCartElement(ctx context.Context, req *pb.AddCartElementRequest) (*pb.AddCartElementResponse, error) {
	_, resp, err := s.addCartElement.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AddCartElementResponse), nil
}

func decodeGRPCAddCartElementRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AddCartElementRequest)
	return AddCartElementRequest{
		Cart_id: int(req.CartId), 
		Item_id: int(req.ItemId), 
		Quantity: req.Quantity,
	}, nil
}

func encodeGRPCAddCartElementResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(AddCartElementResponse)
	return &pb.AddCartElementResponse{
		Err: err2str(resp.Err),
	}, nil
}

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
