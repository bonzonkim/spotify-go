package networks

import (
	"spotify-go/config"
	"spotify-go/spotify"
)

type Router struct {
	*Network
	SpotifyRouter *spotify.Router
}

func newRouter(n *Network, c *config.Config)  {
	sRouter := &spotify.Router{
		Config: c,
	}
	r := &Router {
		Network:	n,
		SpotifyRouter: sRouter,
	}

	n.Router("GET", "/test", r.SpotifyRouter.GetSpotifyToken())
}

