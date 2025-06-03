package handlers

import (
	"log"
	"net/http"
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

//func (h *Handler) SpotifyPageHandler(c *gin.Context) {
//	templates.SpotifyPage().Render(c.Request.Context(), c.Writer)
//}

func (h *Handler) SpotifyPageHandler(c *gin.Context) {
	token, err := c.Cookie("spotify-token")
	log.Printf("token: %v", token)
	if err != nil {
		log.Println("Spotify token is not set")
		// Redirect to login if token is missing
		c.Redirect(http.StatusFound, "/api/auth")
		return
	}

	profile, err := h.Service.SpotifyService.GetUserProfile(token)
	log.Printf("User profile: %v\n", profile)
	if err != nil {
		log.Printf("Error fetching User profile %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Render the SpotifyPage template with the profile data
	templates.SpotifyPage(profile).Render(c.Request.Context(), c.Writer)
}
