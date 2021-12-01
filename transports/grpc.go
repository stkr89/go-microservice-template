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
			decodeGRPCRequest,
			encodeGRPCResponse,
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

func decodeGRPCRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeGRPCResponse(_ context.Context, resp interface{}) (interface{}, error) {
	return resp, nil
}
