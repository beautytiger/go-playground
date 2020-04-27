package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var language string

	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name:        "lang",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &language,
		},
	}

	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Version = "1.2.3"
	app.Action = func(c *cli.Context) error {
		// flag 在命令行中需要放到args的前面
		fmt.Printf("boom! I say! args: %q, flag-lang: %v\n", c.Args(), language)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
