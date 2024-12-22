package networks

import (
	"net/http"
	"spotify-go/config"
	"spotify-go/spotify"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*Network
	SpotifyRouter *spotify.SpotifyRouter
}

func NewRouter(n *Network, c *config.Config) {
	sr := &spotify.SpotifyRouter{
		Config: c,
	}
	r := &Router{
		Network:       n,
		SpotifyRouter: sr,
	}

	n.Router(GET, "/auth", func(ctx *gin.Context) {
		r.SpotifyRouter.GetAuthorization(c, ctx)
	}) 
	n.Router(GET, "/callback", func(ctx *gin.Context) {
		code, err := r.SpotifyRouter.GetCode(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		token, err := r.SpotifyRouter.GetSpotifyToken(code, c)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, token)
	}) 

}
