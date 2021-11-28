package service

import (
	"github.com/shopr-org/grpc-service-template/config"
	"gorm.io/gorm"
)

//go:generate mockery --name=MathDao --output=../mock/
type MathDao interface {
}

type MathDaoImpl struct {
	db gorm.DB
}

func NewMathDaoImpl() *MathDaoImpl {
	return &MathDaoImpl{
		db: config.NewDB(),
	}
}

func NewMathDaoImplArgs(db gorm.DB) MathDaoImpl {
	return MathDaoImpl{
		db: db,
	}
}
