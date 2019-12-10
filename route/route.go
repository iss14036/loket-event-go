package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"loket-event-go/config"
	"loket-event-go/controller"
	"net/http"
)


func Router(c *config.Config) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/customers", controller.GetAllCustomer).Methods("GET")

	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
