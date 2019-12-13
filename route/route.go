package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"loket-event-go/config"
	"loket-event-go/handler"
	"loket-event-go/service"
	"net/http"
)


func Router(c *config.Config, services service.Services) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	customerHandler := handler.NewCustomerHandler(services.NewCustomerHandlerDependencies().CustomerService)
	r.HandleFunc("/customers", customerHandler.GetAllCustomer).Methods(http.MethodGet)

	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
