package database

import (
	"e-commerce-app/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.Customer{})
}
