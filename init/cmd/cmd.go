package cmd

import (
	"spotify-go/config"
	"spotify-go/handlers"
	"spotify-go/networks"
	"spotify-go/service"
)

type Cmd struct {
	config  *config.Config
	network *networks.Network
}

func NewCmd(dirname string) *Cmd {
	c := config.NewConfig(dirname)
	s := service.NewService(c)
	h := handlers.NewHandler(s)
	n := networks.NewNetwork(h)

	n.ServerStart(c.Port)

	return &Cmd{
		config:  c,
		network: n,
	}
}
