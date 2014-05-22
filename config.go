package main

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Database string
	Address  string
}

func newConfig() *config {
	c := &config{
		Database: "stuff.dat",
		Address:  ":9753",
	}
	envconfig.Process("pierolog", c)
	return c
}
