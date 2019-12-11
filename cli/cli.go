package cli

import (
	"context"
	"fmt"
	"github.com/urfave/negroni"
	"log"
	"loket-event-go/config"
	"loket-event-go/connection"
	"loket-event-go/route"
	"loket-event-go/service"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type (
	Cli struct {
		Config *config.Config
	}
	System interface {
		Run()
	}
)

func NewCli() *Cli {
	return &Cli{Config: config.NewConfig()}
}

var wait time.Duration

func (c *Cli) Run() {
	d := make(chan os.Signal, 1)
	signal.Notify(d, os.Interrupt)

	db := connection.GetConnection()
	log.Println("connection for db ",db)

	r := route.Router(c.Config, service.NewServices(db))
	n := negroni.Classic()
	n.UseHandler(r)

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%v",
			c.Config.AppHost,
			c.Config.AppPort),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      n,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	<- d

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
	log.Println("shutting down")
	os.Exit(0)
}