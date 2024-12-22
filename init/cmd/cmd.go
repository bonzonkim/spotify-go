package cmd

import (
	"spotify-go/config"
	"spotify-go/networks"
)

type Cmd struct {
	config  *config.Config
	network *networks.Network
}

func NewCmd() *Cmd {
	c := config.NewConfig()
	n := networks.NewNetwork()

	networks.NewRouter(n, c)

	n.ServerStart(c.Port)

	return &Cmd{
		config:  c,
		network: n,
	}
}
