package cmd

import (
	"spotify-go/config"
	"spotify-go/networks"
)


type Cmd struct {
	config *config.Config
	network *networks.Network
}

func NewCmd() *Cmd {
	c := &Cmd{
		config: config.NewConfig(),
		network: networks.NewNetwork(),
	}

	c.network.ServerStart(c.config.Port)

	return c
}
