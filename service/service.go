package service

import (
	"context"

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

// ProviderMathServiceImpl func initializes a MathServiceImpl
func ProviderMathServiceImpl(logger log.Logger, mathDao MathDao) MathService {
	return &MathServiceImpl{
		logger:  logger,
		mathDao: mathDao,
	}
}

func (s MathServiceImpl) Add(ctx context.Context, numA, numB float32) (float32, error) {
	return numA + numB, nil
}
