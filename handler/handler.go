package handler

import (
	"e-commerce-app/service"

	"go.uber.org/zap"
)

type AllHandler struct {
	CustomerHandler CustomerHadler
}

func NewAllHandler(service service.AllService, log *zap.Logger) AllHandler {
	return AllHandler{
		CustomerHandler: NewCustomerService(service, log),
	}
}
