package service

import "loket-event-go/model"

type (
	CustomerInterface interface {
		GetAllCustomer() []model.Customer
	}
)
