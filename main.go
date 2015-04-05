package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gommbanks"
	app.Usage = "Get latest exchange rates from different banks in Myanmar."
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ye Myat Kaung (Maverick)",
			Email: "mavjs01@gmail.com"},
	}
	app.Commands = []cli.Command{
		{
			Name:    "bank",
			Aliases: []string{"b"},
			Usage:   "bank to get exchange rate from. (currently supported CBM)",
			Action: func(c *cli.Context) {
				switch c.Args().First() {
				case "cbm", "CBM":
					fmt.Println(cbm())
				default:
					fmt.Println(cbm())
				}
			},
		},
	}

	app.Run(os.Args)
}
