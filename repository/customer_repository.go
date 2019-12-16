package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"loket-event-go/domain"
)

type CustomerRepository interface {
	GetAllCustomer(db *sqlx.DB) []domain.Customer
	GetCustomer(db *sqlx.DB, id int64) domain.Customer
	CreateCustomer(db *sqlx.DB, customer domain.Customer) domain.Customer
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

func (repo customerRepository) GetCustomer(db *sqlx.DB, id int64) domain.Customer {
	var customer domain.Customer
	err := db.Get(&customer, "SELECT * FROM CUSTOMERS WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	return customer
}

func (repo customerRepository) CreateCustomer(db *sqlx.DB, customer domain.Customer) domain.Customer {
	tx := db.MustBegin()
	value, err := tx.NamedExec("INSERT INTO CUSTOMERS(name, email, phone) VALUES (:name, :email, :phone)", &customer)
	if err != nil {
		panic(err)
	}
	tx.Commit()
	id, err := value.LastInsertId()
	fmt.Println(id)
	if err != nil {
		panic(err)
	}

	return repo.GetCustomer(db, id)
}