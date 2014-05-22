package server

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

var (
	envPort         = os.Getenv("PORT")
	envRedistogoURL = os.Getenv("REDISTOGO_URL")
)

type config struct {
	Database string
	Port     string

	HipChatAuthToken string
	HipChatRoomID    string
	HipChatFrom      string

	RedisURL      string
	RedisPassword string
}

func newConfig() *config {
	port := "9753"
	if envPort != "" {
		port = envPort
	}

	redisURL := ":6379"
	if envRedistogoURL != "" {
		redisURL = envRedistogoURL
	}

	c := &config{
		Database: "stuff.dat",
		Port:     port,
		RedisURL: redisURL,
	}
	envconfig.Process("pierolog", c)
	return c
}
