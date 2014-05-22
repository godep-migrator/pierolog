package server

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

var (
	envPort = os.Getenv("PORT")
)

type config struct {
	Database string
	Port     string

	HipChatAuthToken string
	HipChatRoomID    string
	HipChatFrom      string
}

func newConfig() *config {
	port := "9753"
	if envPort != "" {
		port = envPort
	}

	c := &config{
		Database: "stuff.dat",
		Port:     port,
	}
	envconfig.Process("pierolog", c)
	return c
}
