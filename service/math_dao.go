package service

import (
	"github.com/go-kit/log"
	"github.com/stkr89/mathsvc/common"
	"github.com/stkr89/mathsvc/config"
	"gorm.io/gorm"
)

//go:generate mocker --dst ../mock/math_dao_mock.go --pkg mock math_dao.go MathDao
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
