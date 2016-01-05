package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func GetApp() *cli.App {
	app := cli.NewApp()
	app.Name = "gowc"
	app.Author = "ariarijp"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "l",
			Usage: "The number of lines in each input file is written to the standard output.",
		},
	}

	app.Action = func(c *cli.Context) {
		l := c.Bool("l")

		if l == true {
			words := countLines(os.Stdin)
			fmt.Println(words)
		} else {
			lines := countWordsInLines(os.Stdin)
			fmt.Println(lines)
		}
	}

	return app
}
