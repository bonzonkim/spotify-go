package networks

import (
	"spotify-go/handlers"

	"github.com/gin-gonic/gin"
)

type Network struct {
	engine *gin.Engine
}

func NewNetwork(h *handlers.Handler) *Network {
	n := &Network{
		engine: gin.New(),
	}

	n.registerGET("/static/*filepath", h.CssHandler)
	n.registerGET("/home", h.HomeHandler)
	n.registerGET("/spotify", h.SpotifyPageHandler)

	NewSpotifyRouter(n, h.Spotify)

	return n
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}
