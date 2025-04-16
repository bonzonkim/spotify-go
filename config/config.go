package config

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
)

type Config struct {
	Port         string
	ClientID     string
	ClientSecret string
}

func NewConfig(dirname string) *Config {
	c, err := LoadConfig(dirname)
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
