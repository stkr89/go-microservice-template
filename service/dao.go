package service

import (
	"github.com/go-kit/log"
	"github.com/shopr-org/grpc-service-template/common"
	"github.com/shopr-org/grpc-service-template/config"
	"gorm.io/gorm"
)

//go:generate mockery --name=MathDao --output=../mock/
type MathDao interface {
}

type MathDaoImpl struct {
	logger log.Logger
	db     gorm.DB
}

func NewMathDaoImpl() *MathDaoImpl {
	return &MathDaoImpl{
		logger: common.NewLogger(),
		db:     config.NewDB(),
	}
}

func NewMathDaoImplArgs(db gorm.DB) MathDaoImpl {
	return MathDaoImpl{
		db: db,
	}
}
