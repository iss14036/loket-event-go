package service

import (
	"github.com/jmoiron/sqlx"
	"loket-event-go/repository"
)

type Services interface {
	NewCustomerHandlerDependencies() CustomerHandlerDependencies
}

type services struct {
	CustomerService
}

func NewServices(db *sqlx.DB) Services {
	cRepository := repository.NewCustomerRepository()
	cService := NewCustomerService(db, cRepository)
	return &services{CustomerService: cService}
}

type CustomerHandlerDependencies struct {
	CustomerService CustomerService
}

func (service services) NewCustomerHandlerDependencies() CustomerHandlerDependencies {
	return CustomerHandlerDependencies{
		CustomerService: service.CustomerService,
	}
}
