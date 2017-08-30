package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	version  string
	revision string
)

func makeApp() *cli.App {
	app := cli.NewApp()
	app.Name = "csv2json"
	app.Usage = "Convert CSV data to JSON"
	app.Version = fmt.Sprintf("%s [%s]", version, revision)

	app.Action = actionMain

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "format, f",
			Usage: "Load JSON format from `FILE`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "print a sample format",
			Action:  actionInit,
		},
	}

	return app
}

func main() {
	app := makeApp()
	app.Run(os.Args)
}
