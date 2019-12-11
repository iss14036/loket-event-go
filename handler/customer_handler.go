package handler

import (
	"github.com/jmoiron/sqlx"
	"loket-event-go/connection"
	"loket-event-go/service"
	"loket-event-go/util"
	"net/http"
)

type customerHandler struct {
	db              *sqlx.DB
	customerService service.CustomerService
}

func NewCustomerHandler(db *sqlx.DB, customerService service.CustomerService) *customerHandler {
	return &customerHandler{
		db: connection.GetConnection(),
		customerService: customerService,
	}
}

func (h *customerHandler) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := h.customerService.GetAllCustomer()
	resp := util.Message(true, "success")
	resp["data"] = customers
	util.Respond(w, resp)
}