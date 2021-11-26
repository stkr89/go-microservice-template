package transport

import (
	"context"

	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/shopr-org/grpc-service-template/endpoints"
	"github.com/shopr-org/grpc-service-template/pb"
)

type gRPCServer struct {
	add gt.Handler
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.MathServiceServer {
	return &gRPCServer{
		add: gt.NewServer(
			endpoints.Add,
			decodeMathRequest,
			encodeMathResponse,
		),
	}
}

func (s *gRPCServer) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func decodeMathRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.MathRequest)
	return endpoints.MathReq{NumA: req.NumA, NumB: req.NumB}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.MathResp)
	return &pb.MathResponse{Result: resp.Result}, nil
}
