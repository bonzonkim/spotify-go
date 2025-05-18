package networks

import (
	"spotify-go/handlers"
	"sync"
)

var (
	spotifyRouterInit     sync.Once
	spotifyRouterInstance *SpotifyRouter
)

type SpotifyRouter struct {
	router  *Network
	handler *handlers.SpotifyHandler
}

func NewSpotifyRouter(n *Network, h *handlers.SpotifyHandler) *SpotifyRouter {
	spotifyRouterInit.Do(func() {
		spotifyRouterInstance = &SpotifyRouter{
			router:  n,
			handler: h,
		}
		api := n.engine.Group("/api")
		api.GET("/auth", h.GetAuthorization)
		api.GET("/callback", h.Callback)
	})
	return spotifyRouterInstance
}
