package main

import (
	"log"
	"os"

	"github.com/nomura-lab/dps/api"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dps"
	app.Usage = "dps shows processes in each containers"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "apiVersion, apiversion",
			Value: "default",
			Usage: "API version",
		},
	}

	app.Action = func(c *cli.Context) error {
		api.RunDps(c.String("apiVersion"))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
