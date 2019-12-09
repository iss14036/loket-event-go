package model

import (
	"loket-event-go/config"
	"loket-event-go/connection"
	"time"
)

type Customer struct {
	Id        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt  time.Time `json:"update_at" db:"updated_at"`
}

func GetAllCustomer() []Customer {
	db := connection.GetConnection(config.NewConfig())
	var customers []Customer
	err := db.Select(&customers, "SELECT * FROM customers order by id ASC")
	if err != nil {
		panic(err.Error())
	}
	return customers
}