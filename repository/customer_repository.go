package repository

import (
	"github.com/jmoiron/sqlx"
	"loket-event-go/domain"
)

type CustomerRepository interface {
	GetAllCustomer(db *sqlx.DB) []domain.Customer
}

type customerRepository struct {
}

func NewCustomerRepository() *customerRepository {
	return &customerRepository{}
}

func (repo customerRepository) GetAllCustomer(db *sqlx.DB) []domain.Customer {
	var customers []domain.Customer
	err := db.Select(&customers, "SELECT * FROM customers order by id ASC")
	if err != nil {
		panic(err.Error())
	}
	return customers
}
