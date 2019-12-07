package controller

import (
	"loket-event-go/config"
	"loket-event-go/connection"
	"loket-event-go/model"
	"loket-event-go/util"
	"net/http"
)

var GetAllCustomer = func(w http.ResponseWriter, r *http.Request) {
	db := connection.GetConnection(config.NewConfig())
	var customers []model.Customer
	err := db.Select(&customers, "SELECT * FROM customers order by id ASC")
	if err != nil {
		panic(err.Error())
	}
	resp := util.Message(true, "success")
	resp["data"] = customers
	util.Respond(w, resp)
}
