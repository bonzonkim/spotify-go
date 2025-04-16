package cmd

import (
	"spotify-go/config"
	"spotify-go/networks"
)

type Cmd struct {
	config  *config.Config
	network *networks.Network
}

func NewCmd(dirname string) *Cmd {
	c := config.NewConfig(dirname)
	n := networks.NewNetwork()

	networks.NewRouter(n, c)

	n.ServerStart(c.Port)

	return &Cmd{
		config:  c,
		network: n,
	}
}
