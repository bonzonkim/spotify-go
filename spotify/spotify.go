package spotify

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"spotify-go/config"
)

type Router struct {
	Config *config.Config
}

type SpotifyToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

const (
	RedirectURI = "localhost:8080/callback"
)

func (s *Router)GetSpotifyToken(code string, c *config.Config) (*SpotifyToken, error) {
	basicToken := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.ClientID, c.ClientSecret)))
	authHeader := "Basic " + basicToken
	data := url.Values{}
	data.Set("grant_type", "authorized_code")
	data.Set("redirect_uri", RedirectURI)
	data.Set("code", code)

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch token" + resp.Status)
	}
	var token SpotifyToken
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}
	log.Println(&token)
	return &token, nil
}
