package service

type CustomerResponse struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
}

type CuxtomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer() (*CustomerResponse, error)
}
