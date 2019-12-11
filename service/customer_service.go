package service

import (
	"github.com/jmoiron/sqlx"
	"loket-event-go/connection"
	"loket-event-go/domain"
	"loket-event-go/repository"
)

type (
	CustomerService interface {
		GetAllCustomer() []domain.Customer
	}
	customerService struct {
		DB                 *sqlx.DB
		customerRepository repository.CustomerRepository
	}
)

func NewCustomerService(DB *sqlx.DB, repository repository.CustomerRepository) *customerService {
	return &customerService{
		DB:                 connection.GetConnection(),
		customerRepository: repository,
	}
}

func (cs *customerService) GetAllCustomer() []domain.Customer {
	customers := cs.customerRepository.GetAllCustomer(cs.DB)
	return customers
}