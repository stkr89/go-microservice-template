package service

import (
	"github.com/stkr89/mathsvc/common"
	"github.com/stkr89/mathsvc/types"

	"github.com/go-kit/kit/log"
)

// MathService interface
type MathService interface {
	Add(request *types.MathRequest) (*types.MathResponse, error)
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

func (s MathServiceImpl) Add(request *types.MathRequest) (*types.MathResponse, error) {
	return nil, nil
}
