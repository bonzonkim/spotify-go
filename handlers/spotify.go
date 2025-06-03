package handlers

import (
	"log"
	"net/http"
	"spotify-go/internal/templates"
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
	//https://accounts.spotify.com/authorize?response_type=code&client_id=74fce79bb47b4576a2057a92ac976cac&scope=user-top-read+user-read-email&redirect_uri=http://localhost:8080/api/callback
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
	//SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)
	c.SetCookie("spotify-token", token.AccessToken, token.ExpiresIn, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/spotify")

	//profile, err := h.service.GetUserProfile(token)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//c.JSON(http.StatusOK, profile)
}

func (h *SpotifyHandler) UserProfile(c *gin.Context) {
	token, err := c.Cookie("spotify-token")
	log.Printf("token: %v", token)
	if err != nil {
		log.Println("Spotify token is not set")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Spotify token is missing"})
		return
	}

	profile, err := h.service.GetUserProfile(token)
	log.Printf("User profile: %v\n", profile)
	if err != nil {
		log.Printf("Error fetching User profile %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
	//c.HTML(http.StatusOK, "SpotifyUserProfile.templ", profile)
	templates.SpotifyPage(profile).Render(c.Request.Context(), c.Writer)
}
