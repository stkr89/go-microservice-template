package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

// MathService interface
type MathService interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
}

type mathService struct {
	logger  log.Logger
	mathDao MathDao
}

// NewService func initializes a mathService
func NewService(logger log.Logger, mathDao MathDao) MathService {
	return &mathService{
		logger:  logger,
		mathDao: mathDao,
	}
}

func (s mathService) Add(ctx context.Context, numA, numB float32) (float32, error) {
	return numA + numB, nil
}
