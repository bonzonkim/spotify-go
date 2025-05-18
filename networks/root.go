package networks

import (
	"path/filepath"
	"spotify-go/handlers"
	"spotify-go/internal/templates"

	"github.com/gin-gonic/gin"
)

type Network struct {
	engine *gin.Engine
}

func NewNetwork(h *handlers.Handler) *Network {
	n := &Network{
		engine: gin.New(),
	}

	NewSpotifyRouter(n, h.Spotify)

	n.registerGET("/static/*filepath", func(c *gin.Context) {
		path := filepath.Clean(c.Param("filepath"))

		if filepath.Ext(path) == ".css" {
			c.Header("Content-type", "text/css; charset=utf-8")
		}
		c.File("./static" + path)
	})

	n.registerGET("/home", func(c *gin.Context) {
		templates.Home().Render(c.Request.Context(), c.Writer)
	})
	n.registerGET("/spotify", func(c *gin.Context) {
		templates.SpotifyPage().Render(c.Request.Context(), c.Writer)
	})
	return n
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}
