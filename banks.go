package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func cbm() string {
	// parse json
	type currency struct {
		USD string
		EUR string
	}

	type jsonUser struct {
		Rates currency
	}

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

	user := jsonUser{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("USD: %s\nEUR: %s", user.Rates.USD, user.Rates.EUR)
}
