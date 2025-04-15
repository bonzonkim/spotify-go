package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	ClientID     string
	ClientSecret string
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load .env file:" + err.Error())
	}

	c := &Config{
		Port:         GetEnv("SERVERPORT"),
		ClientID:     GetEnv("CLIENTID"),
		ClientSecret: GetEnv("CLIENTSECRET"),
	}

	return c
}

func GetEnv(key string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	panic("Missing required .env variable: " + key)
}
