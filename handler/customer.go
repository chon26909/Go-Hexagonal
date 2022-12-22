package handler

import (
	"hexagonal/service"
	"net/http"
)

type customerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) customerHandler {
	return customerHandler{customerService: customerService}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	println("Hwllo World")
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {}
