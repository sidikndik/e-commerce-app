package repository

import (
	"e-commerce-app/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewCustomerRepository(db *gorm.DB, log *zap.Logger) CustomerRepository {
	return CustomerRepository{
		DB:     db,
		Logger: log,
	}
}

func (customerRepo *CustomerRepository) Create(customer *model.Customer) error {
	query := "INSET INTO customers (name, email, phone, password) VALUES ($1, $2, $3, $4) RETURNING id"
	err := customerRepo.DB.Raw(query, customer.Name, customer.Email, customer.Phone, customer.Password).Scan(&customer.ID).Error
	return err
}
