package handlers

import (
	"path/filepath"
	"spotify-go/internal/templates"
	"spotify-go/service"
	"sync"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
	Spotify *SpotifyHandler
}

var (
	handlerInit     sync.Once
	handlerInstance *Handler
)

func NewHandler(s *service.Service) *Handler {
	handlerInit.Do(func() {
		handlerInstance = &Handler{
			Service: s,
		}
		handlerInstance.Spotify = NewSpotifyHandler(s.SpotifyService)
	})

	return handlerInstance
}

func (h *Handler) CssHandler(c *gin.Context) {
	path := filepath.Clean(c.Param("filepath"))

	if filepath.Ext(path) == ".css" {
		c.Header("Content-type", "text/css; charset=utf-8")
	}
	c.File("./static" + path)
}

func (h *Handler) HomeHandler(c *gin.Context) {
	templates.Home().Render(c.Request.Context(), c.Writer)
}

func (h *Handler) SpotifyPageHandler(c *gin.Context) {
	templates.SpotifyPage().Render(c.Request.Context(), c.Writer)
}
