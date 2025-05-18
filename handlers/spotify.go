package handlers

import (
	"net/http"
	"spotify-go/service"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	spotifyHandlerInit     sync.Once
	spotifyHandlerInstance *SpotifyHandler
)

type SpotifyHandler struct {
	service *service.SpotifyService
}

func NewSpotifyHandler(s *service.SpotifyService) *SpotifyHandler {
	spotifyHandlerInit.Do(func() {
		spotifyHandlerInstance = &SpotifyHandler{
			service: s,
		}
	})
	return spotifyHandlerInstance
}

func (h *SpotifyHandler) GetAuthorization(c *gin.Context) {
	authURL := h.service.GetAutorizationURL()
	c.Redirect(http.StatusFound, authURL)
}

func (h *SpotifyHandler) Callback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorization code not found."})
		return
	}

	token, err := h.service.GetSpotifyToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.service.GetUserProfile(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)

}
