package repository

import (
	"database/sql"
	"e-commerce-app/model"

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

func (customerRepo *CustomerRepository) Create(customer *model.Customer) error {
	query := "INSET INTO customers (name, email, phone, password) VALUES ($1, $2, $3, $4) RETURNING id"
	err := customerRepo.DB.QueryRow(query, customer.Name, customer.Email, customer.Phone, customer.Password).Scan(&customer.ID)
	return err
}
