package service

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
	types "spotify-go/types/spotify"
	"sync"
)

const BASE_URL = "https://api.spotify.com/v1/me"

var (
	spotifyServiceInit     sync.Once
	spotifyServiceInstance *SpotifyService
)

type SpotifyService struct {
	config *config.Config
}

type SpotifyToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

const (
	RedirectURI = "http://localhost:8080/api/callback"
	Scope       = "user-top-read+user-read-email"
)

func NewSpotifyService(c *config.Config) *SpotifyService {
	spotifyServiceInit.Do(func() {
		spotifyServiceInstance = &SpotifyService{
			config: c,
		}
	})
	return spotifyServiceInstance
}

func (s *SpotifyService) GetAutorizationURL() string {
	return fmt.Sprintf(
		"https://accounts.spotify.com/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s",
		s.config.ClientID,
		Scope,
		RedirectURI,
	)
}

func (s *SpotifyService) GetSpotifyToken(code string) (*SpotifyToken, error) {
	basicToken := base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf(
			"%s:%s",
			s.config.ClientID,
			s.config.ClientSecret,
		)),
	)

	authHeader := "Basic " + basicToken
	params := url.Values{}
	params.Set("grant_type", "authorization_code")
	params.Set("redirect_uri", RedirectURI)
	params.Set("code", code)

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(params.Encode()))
	if err != nil {
		log.Printf("Failed to make new request %v", err)
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
		return nil, errors.New("Failed to fetch token" + resp.Status)
	}

	var token SpotifyToken
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *SpotifyService) GetUserProfile(accessToken string) (types.SpotifyUser, error) {

	req, err := http.NewRequest("GET", BASE_URL, nil)
	if err != nil {
		return types.SpotifyUser{}, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return types.SpotifyUser{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return types.SpotifyUser{}, fmt.Errorf("failed to fetch user's data %v", resp.Status)
	}

	var userProfile types.SpotifyUser
	if err := json.NewDecoder(resp.Body).Decode(&userProfile); err != nil {
		return types.SpotifyUser{}, fmt.Errorf("failed to decode response body %v", err)
	}
	log.Printf("userProfile %v", userProfile)
	return userProfile, nil
}

func (s *SpotifyService) GetUsersTopTracks(accessToken string) (types.SpotifyTopTracks, error) {
	req, err := http.NewRequest(http.MethodGet, BASE_URL+"/top/tracks", nil)
	if err != nil {
		log.Printf("Failed to create request %v", err)
		return types.SpotifyTopTracks{}, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to Request %v", err)
		return types.SpotifyTopTracks{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return types.SpotifyTopTracks{}, fmt.Errorf("Failed to fetch user's top tracks data %v", resp.Status)
	}

	var userTopTracks types.SpotifyTopTracks
	if err := json.NewDecoder(resp.Body).Decode(&userTopTracks); err != nil {
		return types.SpotifyTopTracks{}, fmt.Errorf("Failed to decode response body %v", err)
	}
	log.Printf("User top track data %v", userTopTracks)
	return userTopTracks, nil
}
