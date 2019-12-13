package service

import (
	"loket-event-go/connection"
	"loket-event-go/repository"
)

type Services interface {
	NewCustomerHandlerDependencies() CustomerHandlerDependencies
}

type services struct {
	CustomerService
}

func NewServices(db connection.Mysql) Services {
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
