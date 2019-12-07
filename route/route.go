package route

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"loket-event-go/config"
	"loket-event-go/connection"
	"net/http"
	"os"
	"time"
)

var wait time.Duration

func Run(c *config.Config, s <-chan os.Signal) {
	db := connection.GetConnection(c)
	fmt.Sprintf("connection for db ",db)

	r := mux.NewRouter()

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%v",
			c.AppHost,
			c.AppPort),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	r.HandleFunc("/", index)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	<- s

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
