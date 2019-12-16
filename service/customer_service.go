package service

import (
	"loket-event-go/connection"
	"loket-event-go/domain"
	"loket-event-go/repository"
)

type (
	CustomerService interface {
		GetAllCustomer() []domain.Customer
		CreateCustomer(customer domain.Customer) domain.Customer
		GetCustomer(id int64) domain.Customer
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

func (cs *customerService) CreateCustomer(customer domain.Customer) domain.Customer {
	newCustomer := cs.customerRepository.CreateCustomer(cs.DB.GetConnection(), customer)
	return newCustomer
}

func (cs *customerService) GetCustomer(id int64) domain.Customer {
	customer := cs.customerRepository.GetCustomer(cs.DB.GetConnection(), id)
	return customer
}