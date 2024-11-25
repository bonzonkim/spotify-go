package spotify

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"spotify-go/config"
)

type SpotifyToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

const (
	RedirectURI = "localhost:8080/callback"
)

func getSpotifyToken(code string, c *config.Config) (*SpotifyToken, error) {
	data := url.Values{}
	data.Set("grant_type", "authorized_code")
	data.Set("redirect_uri", RedirectURI)
	data.Set("code", code)

	basicToken := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.ClientID, c.ClientSecret)))
}
