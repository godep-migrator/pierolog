package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/meatballhat/pierolog/server"
)

func main() {
	app := cli.NewApp()
	app.Name = "pierolog"
	app.Usage = "make benefit pierogi travelogue"
	app.Commands = []cli.Command{
		{
			Name:      "serve",
			ShortName: "s",
			Usage:     "serve the pierolog!",
			Action: func(c *cli.Context) {
				server.Main()
			},
		},
	}
	app.Run(os.Args)
}
