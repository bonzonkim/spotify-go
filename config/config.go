package config

import (
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	ClientID     string
	ClientSecret string
}

func NewConfig(dirname string) *Config {
	c, err := LoadConfig(dirname)
	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load .env: " + err.Error())
	}
	c.ClientID = os.Getenv("CLIENTID")
	c.ClientSecret = os.Getenv("CLIENTSECRET")
	fmt.Printf("%+v", c)
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	return c
}

func LoadConfig(dirname string) (*Config, error) {
	cueConfig := &load.Config{
		Dir: dirname,
	}

	buildInstances := load.Instances([]string{}, cueConfig)
	runtimeInstances := cue.Build(buildInstances)
	instance := runtimeInstances[0]

	var config Config
	if err := instance.Value().Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
