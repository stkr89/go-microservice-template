package endpoints

import (
	"context"
	"github.com/stkr89/mathsvc/types"

	"github.com/go-kit/kit/endpoint"
	"github.com/stkr89/mathsvc/service"
)

type Endpoints struct {
	Add endpoint.Endpoint
}

func MakeEndpoints(s service.MathService) Endpoints {
	return Endpoints{
		Add: makeAddEndpoint(s),
	}
}

func makeAddEndpoint(s service.MathService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*types.MathRequest)
		return s.Add(req)
	}
}
