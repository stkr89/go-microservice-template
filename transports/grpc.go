package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/stkr89/go-microservice-template/middleware"
	"github.com/stkr89/go-microservice-template/types"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/stkr89/go-microservice-template/endpoints"
	"github.com/stkr89/go-microservice-template/pb"
)

type gRPCServer struct {
	add gt.Handler
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoints.Endpoints) pb.MathServiceServer {
	return &gRPCServer{
		add: gt.NewServer(
			endpoint.Chain(
				middleware.ValidateAddInput(),
				middleware.ConformAddInput(),
			)(endpoints.Add),
			decodeAddGRPCRequest,
			encodeAddGRPCResponse,
		),
	}
}

func (s gRPCServer) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func decodeAddGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	reqpb := request.(*pb.MathRequest)
	b, err := json.Marshal(reqpb)
	if err != nil {
		return nil, err
	}

	var req types.MathRequest
	err = json.Unmarshal(b, &req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func encodeAddGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	respb := response.(*types.MathResponse)
	b, err := json.Marshal(respb)
	if err != nil {
		return nil, err
	}

	var resp pb.MathResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
