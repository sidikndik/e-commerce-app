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
	query := "INSERT INTO customers (name, email, phone, password) VALUES ($1, $2, $3, $4) RETURNING id"
	err := customerRepo.DB.Raw(query, customer.Name, customer.Email, customer.Phone, customer.Password).Scan(&customer.ID).Error
	return err
}

func (customerRepo *CustomerRepository) GetAll() (*[]model.Customer, error) {
	var customers []model.Customer
	query := "SELECT name, email, phone, password FROM customers"
	rows, err := customerRepo.DB.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.Customer
		if err := customerRepo.DB.ScanRows(rows, &customers); err != nil {
			return nil, err
		}
		customers = append(customers, user)
	}
	return &customers, nil
}
