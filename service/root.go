package service

import (
	"spotify-go/config"
	"sync"
)

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	Config         *config.Config
	SpotifyService *SpotifyService
}

func NewService(c *config.Config) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			Config: c,
		}
		serviceInstance.SpotifyService = NewSpotifyService(c)
	})
	return serviceInstance
}
