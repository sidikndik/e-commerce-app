package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AllRepository struct {
	CustomerRep CustomerRepository
}

func NewAllRepository(db *gorm.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		CustomerRep: NewCustomerRepository(db, log),
	}
}
