package service

import (
	"hexagonal/repository"
	"log"
)

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) customerService {
	return customerService{customerRepository: customerRepository}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerRepository.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := []CustomerResponse{}
	for _, customer := range customers {
		row := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		res = append(res, row)
	}

	return res, nil

}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.customerRepository.GetById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &customerResponse, nil
}
