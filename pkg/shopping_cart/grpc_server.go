package shopping_cart

import (
	"context"
	"errors"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	"shopping_cart/pb"
)

type grpcServer struct {
	addCart grpctransport.Handler
	addItem grpctransport.Handler
}

func NewGRPCServer(endpoints Endpoints) pb.CartServer {
	return &grpcServer{
		addCart: grpctransport.NewServer(
			endpoints.AddCartEndpoint,
			decodeGRPCAddCartRequest,
			encodeGRPCAddCartResponse,
		),
		addItem: grpctransport.NewServer(
			endpoints.AddItemEndpoint,
			decodeGRPCAddItemRequest,
			encodeGRPCAddItemResponse,
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
	return AddCartRequest{Id: int(req.Id)}, nil
}

func encodeGRPCAddCartResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(AddCartResponse)
	return &pb.AddCartResponse{Id: int64(resp.Id), Err: err2str(resp.Err)}, nil
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
	return AddItemRequest{Id: int(req.Id), Detail: req.Detail, Price: float64(req.Price)}, nil
}

func encodeGRPCAddItemResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(AddItemResponse)
	return &pb.AddItemResponse{Id: int64(resp.Id), Err: err2str(resp.Err)}, nil
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
