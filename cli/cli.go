package cli

import (
	"loket-event-go/config"
	"loket-event-go/route"
	"os"
	"os/signal"
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

func (c *Cli) Run() {
	d := make(chan os.Signal)
	signal.Notify(d, os.Interrupt)

	route.Run(c.Config, d)
}