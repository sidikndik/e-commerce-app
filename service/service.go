package service

import (
	"e-commerce-app/repository"

	"go.uber.org/zap"
)

type AllService struct {
	CustomerService CustomerService
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		CustomerService: NewCustomerService(repo, log),
	}
}
