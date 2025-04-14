package networks

import (
	"log"
	"net/http"
	"spotify-go/config"
	"spotify-go/spotify"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*Network
	SpotifyRouter *spotify.SpotifyRouter
	Token         *spotify.SpotifyToken
}

func NewRouter(n *Network, c *config.Config) {
	sr := &spotify.SpotifyRouter{}
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

		r.Token = token

		ctx.JSON(http.StatusOK, token)
	})

	n.Router(GET, "/test", func(ctx *gin.Context) {
		if r.Token == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Spotify token not found"})
			return
		}

		result, err := r.SpotifyRouter.GetUserProfile(r.Token)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		log.Printf("result: %v", result)
		ctx.JSON(http.StatusOK, result)
	})
}
