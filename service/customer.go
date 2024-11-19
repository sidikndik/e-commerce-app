package service

import (
	"e-commerce-app/repository"

	"go.uber.org/zap"
)

type CustomerService struct {
	Repo repository.AllRepository
	Log  *zap.Logger
}

func NewCustomerService(repo repository.AllRepository, log *zap.Logger) CustomerService {
	return CustomerService{
		Repo: repo,
		Log:  log,
	}
}
