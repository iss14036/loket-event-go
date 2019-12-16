package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"loket-event-go/config"
	"loket-event-go/handler"
	"loket-event-go/service"
	"net/http"
)


func Router(c *config.Config, services service.Services) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/metric", promhttp.Handler())
	r.HandleFunc("/", index)

	r = CustomerRouter(r, services)

	return r
}

func CustomerRouter(r *mux.Router, services service.Services) *mux.Router {
	customerHandler := handler.NewCustomerHandler(services.NewCustomerHandlerDependencies().CustomerService)
	r.HandleFunc("/customers", customerHandler.GetAllCustomer).Methods(http.MethodGet)
	r.HandleFunc("/customers", customerHandler.CreateCustomer).Methods(http.MethodPost)

	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
