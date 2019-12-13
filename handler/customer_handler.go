package handler

import (
	"loket-event-go/connection"
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