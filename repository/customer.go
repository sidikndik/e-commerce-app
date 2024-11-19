package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type CustomerRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewCustomerRepository(db *sql.DB, log *zap.Logger) CustomerRepository {
	return CustomerRepository{
		DB:     db,
		Logger: log,
	}
}
