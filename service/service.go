package service

import (
	"context"
	"github.com/shopr-org/grpc-service-template/common"

	"github.com/go-kit/kit/log"
)

// MathService interface
type MathService interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
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

func (s MathServiceImpl) Add(ctx context.Context, numA, numB float32) (float32, error) {
	return numA + numB, nil
}
