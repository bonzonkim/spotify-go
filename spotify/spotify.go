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

	"github.com/gin-gonic/gin"
)

type SpotifyRouter struct {
	//Config *config.Config
}

type SpotifyToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

const (
	RedirectURI = "http://localhost:8080/callback"
	Scope       = "user-top-read+user-read-email"
)

func (sr *SpotifyRouter) GetAuthorization(c *config.Config, ctx *gin.Context) {
	authUrl := fmt.Sprintf(
		"https://accounts.spotify.com/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s",
		c.ClientID, Scope, RedirectURI,
	)
	ctx.Redirect(http.StatusFound, authUrl)
}

func (sr *SpotifyRouter) GetCode(ctx *gin.Context) (string, error) {
	code := ctx.Query("code")
	if code == "" {
		return "", errors.New("authorization code not found in callback")
	}
	fmt.Printf("CODE : ", code)
	return code, nil
}

func (sr *SpotifyRouter) GetSpotifyToken(code string, c *config.Config) (*SpotifyToken, error) {
	basicToken := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.ClientID, c.ClientSecret)))
	authHeader := "Basic " + basicToken
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", RedirectURI)
	data.Set("code", code)

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(data.Encode()));
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

func (sr * SpotifyRouter) GetUserProfile(token *SpotifyToken) (map[string]interface{}, error) {
	accessToken := token.AccessToken

	// create http request
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer " + accessToken)

	// create http client to request
	client := &http.Client{}
	// requesting
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user's data %v", resp.Status)
	}

	var userProfile map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userProfile); err != nil {
		return nil, fmt.Errorf("failed to decode reponse body %v", err)
	}
	log.Printf("userProfile %v", userProfile)
	return userProfile, nil
}
