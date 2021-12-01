package service

import (
	"context"
	"github.com/shopr-org/grpc-service-template/common"
	"github.com/shopr-org/grpc-service-template/pb"

	"github.com/go-kit/kit/log"
)

// MathService interface
type MathService interface {
	Add(ctx context.Context, request *pb.MathRequest) (*pb.MathResponse, error)
}

type MathServiceImpl struct {
	logger  log.Logger
	mathDao MathDao
}

func NewMathServiceImpl() *MathServiceImpl {
	return &MathServiceImpl{
		logger:  common.NewLogger(),
		mathDao: NewMathDaoImpl(),
	}
}

func NewMathServiceImplArgs(logger log.Logger, mathDao MathDao) MathService {
	return &MathServiceImpl{
		logger:  logger,
		mathDao: mathDao,
	}
}

func (s MathServiceImpl) Add(ctx context.Context, request *pb.MathRequest) (*pb.MathResponse, error) {
	return nil, nil
}
