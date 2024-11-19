package handler

import (
	"e-commerce-app/service"
	"net/http"

	"go.uber.org/zap"
)

type CustomerHadler struct {
	Service service.AllService
	Log     *zap.Logger
}

func NewCustomerService(service service.AllService, log *zap.Logger) CustomerHadler {
	return CustomerHadler{
		Service: service,
		Log:     log,
	}
}

func (CustomerHadler *CustomerHadler) Create(w http.ResponseWriter, r *http.Request) {

}
