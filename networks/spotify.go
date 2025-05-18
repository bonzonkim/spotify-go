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

//func NewSpotifyRouter(n *Network, h *handlers.SpotifyHandler) *SpotifyRouter {
//	r := &SpotifyRouter{
//		router:  n,
//		handler: h,
//	}
//
//	api := r.router.engine.Group("/api")
//
//	api.GET("/auth", func(ctx *gin.Context) {
//		r.handler.GetAuthorization(ctx)
//	})
//
//	api.GET("/callback", func(ctx *gin.Context) {
//		code, err := r.handler.GetCode(ctx)
//		if err != nil {
//			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
//			return
//		}
//
//		token, err := r.handler.GetSpotifyToken(code)
//		if err != nil {
//			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
//			return
//		}
//
//		r.Token = token
//
//		ctx.JSON(http.StatusOK, token)
//	})
//	api.GET("/test", func(ctx *gin.Context) {
//		if r.Token == nil {
//			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Spotify token not found"})
//			return
//		}
//
//		result, err := r.handler.GetUserProfile(r.Token)
//		if err != nil {
//			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
//			return
//		}
//		log.Printf("result: %v", result)
//		ctx.JSON(http.StatusOK, result)
//	})
//
//	api.GET("/user-profile", func(ctx *gin.Context) {
//		log.Println("hello")
//	})
//
//	//n.Router(GET, "/static/*filepath", func(ctx *gin.Context) {
//	//	path := filepath.Clean(ctx.Param("filepath"))
//	//
//	//	if filepath.Ext(path) == ".css" {
//	//		ctx.Header("Content-Type", "text/css; charset=utf-8")
//	//	}
//	//
//	//	ctx.File("./static/" + path)
//	//})
//	//
//	//n.Router(GET, "/home", func(ctx *gin.Context) {
//	//	templates.Home().Render(ctx.Request.Context(), ctx.Writer)
//	//})
//	//
//	//n.Router(GET, "/spotify", func(ctx *gin.Context) {
//	//	templates.SpotifyPage().Render(ctx.Request.Context(), ctx.Writer)
//	//})
//}
