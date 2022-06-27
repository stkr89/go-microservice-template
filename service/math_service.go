package service

import (
	"github.com/stkr89/go-microservice-template/common"
	"github.com/stkr89/go-microservice-template/dao"
	"github.com/stkr89/go-microservice-template/types"

	"github.com/go-kit/kit/log"
)

// MathService interface
type MathService interface {
	Add(request *types.MathRequest) (*types.MathResponse, error)
}

type MathServiceImpl struct {
	logger  log.Logger
	mathDao dao.MathDao
}

func NewMathServiceImpl() *MathServiceImpl {
	return &MathServiceImpl{
		logger:  common.NewLogger(),
		mathDao: dao.NewMathDaoImpl(),
	}
}

func NewMathServiceImplArgs(logger log.Logger, mathDao dao.MathDao) MathService {
	return &MathServiceImpl{
		logger:  logger,
		mathDao: mathDao,
	}
}

func (s MathServiceImpl) Add(request *types.MathRequest) (*types.MathResponse, error) {
	return nil, nil
}
