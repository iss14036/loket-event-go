package service

import (
	"loket-event-go/connection"
	"loket-event-go/domain"
	"loket-event-go/repository"
)

type (
	CustomerService interface {
		GetAllCustomer() []domain.Customer
	}
	customerService struct {
		DB                 connection.Mysql
		customerRepository repository.CustomerRepository
	}
)

func NewCustomerService(db connection.Mysql, repository repository.CustomerRepository) *customerService {
	return &customerService{
		DB:                 db,
		customerRepository: repository,
	}
}

func (cs *customerService) GetAllCustomer() []domain.Customer {
	customers := cs.customerRepository.GetAllCustomer(cs.DB.GetConnection())
	return customers
}