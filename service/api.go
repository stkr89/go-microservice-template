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
	logger log.Logger
}

// NewService func initializes a mathService
func NewService(logger log.Logger) MathService {
	return &mathService{
		logger: logger,
	}
}

func (s mathService) Add(ctx context.Context, numA, numB float32) (float32, error) {
	return numA + numB, nil
}
