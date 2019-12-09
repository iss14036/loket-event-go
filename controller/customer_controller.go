package controller

import (
	"loket-event-go/model"
	"loket-event-go/util"
	"net/http"
)

var GetAllCustomer = func(w http.ResponseWriter, r *http.Request) {
	customers := model.GetAllCustomer()
	resp := util.Message(true, "success")
	resp["data"] = customers
	util.Respond(w, resp)
}
