package handler

import (
	"e-commerce-app/helper"
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

func (CustomerHadler *CustomerHadler) GetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := CustomerHadler.Service.CustomerService.GetAll()
	if err != nil {
		CustomerHadler.Log.Error("error get all data :", zap.Error(err))
		helper.BadResponse(w, err.Error(), http.StatusBadRequest)
	}

	helper.SuccessResponseWithData(w, "success", http.StatusOK, customers)
}
