package handlers

import (
	"spotify-go/service"
	"sync"
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
