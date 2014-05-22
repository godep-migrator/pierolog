package main

import (
	"fmt"
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
		{
			Name:  "implode",
			Usage: "destroy everything",
			Action: func(c *cli.Context) {
				fmt.Println("ZOMG DESTROYING REDIS")
			},
		},
	}
	app.Run(os.Args)
}
