package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"net/http"
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

func cbm() string {
	var err error

	// request http api
	res, err := http.Get("http://forex.cbm.gov.mm/api/latest")
	if err != nil {
		log.Fatal(err)
	}

	// read body
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// parse json
	type currency struct {
		USD string
		EUR string
	}

	type jsonUser struct {
		Rates currency
	}
	user := jsonUser{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("USD: %s\nEUR: %s", user.Rates.USD, user.Rates.EUR)
}
