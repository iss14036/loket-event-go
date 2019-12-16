package handler

import (
	"encoding/json"
	"fmt"
	"loket-event-go/connection"
	"loket-event-go/domain"
	"loket-event-go/service"
	"loket-event-go/util"
	"net/http"
)

type customerHandler struct {
	db              connection.Mysql
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *customerHandler {
	return &customerHandler{
		customerService: customerService,
	}
}

func (h *customerHandler) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := h.customerService.GetAllCustomer()
	resp := util.Message(true, "success")
	resp["data"] = customers
	util.Respond(w, resp)
}

func (h *customerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer domain.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		panic(err)
	}
	fmt.Println(customer)
	newCustomer := h.customerService.CreateCustomer(customer)
	resp := util.Message(true, "customer created")
	resp["data"] = newCustomer
	util.Respond(w, resp)
}
