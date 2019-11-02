package shopping_cart

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"shopping_cart/pb"
)

func NewGRPCClient(conn *grpc.ClientConn) CartService {

	addCartEndpoint := grpctransport.NewClient(
		conn,
		"pb.CartService",
		"AddCart",
		encodeGRPCAddCartRequest,
		decodeGRPCAddCartResponse,
		pb.AddCartResponse{},
	).Endpoint()

	addItemEndpoint := grpctransport.NewClient(
		conn,
		"pb.CartService",
		"AddItem",
		encodeGRPCAddItemRequest,
		decodeGRPCAddItemResponse,
		pb.AddItemResponse{},
	).Endpoint()

	return Endpoints{
		AddCartEndpoint: addCartEndpoint,
		AddItemEndpoint: addItemEndpoint,
	}
}

func encodeGRPCAddCartRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(AddCartRequest)
	return &pb.AddCartRequest{Id: int64(req.Id)}, nil
}

func decodeGRPCAddCartResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.AddCartResponse)
	return AddCartResponse{Id: int(reply.Id), Err: str2err(reply.Err)}, nil
}

func encodeGRPCAddItemRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(AddItemRequest)
	return &pb.AddItemRequest{Id: int64(req.Id), Detail: req.Detail, Price: float32(req.Price)}, nil
}

func decodeGRPCAddItemResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.AddItemResponse)
	return AddItemResponse{Id: int(reply.Id), Err: str2err(reply.Err)}, nil
}

