package server

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Database string
	Port     string

	HipChatAuthToken string
	HipChatRoomID    string
	HipChatFrom      string
}

func newConfig() *config {
	c := &config{
		Database: "stuff.dat",
		Port:     "9753",
	}
	envconfig.Process("pierolog", c)
	return c
}
