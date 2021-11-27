package service

import "gorm.io/gorm"

//go:generate mockery --name=MathDao --output=../mock/
type MathDao interface {
}

type MathDaoImpl struct {
	db gorm.DB
}

func ProviderMathDaoImpl(db gorm.DB) MathDaoImpl {
	return MathDaoImpl{
		db: db,
	}
}
