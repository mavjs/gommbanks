package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
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
    //    Info string
    //    Description string
    //    Timestamp string
        Rates currency
    }
    user := jsonUser{}

    err = json.Unmarshal(body, &user)
    if err != nil {
        log.Fatal(err)
    }

    // fmt.Println(user.Info)
    // fmt.Println(user.Description)
    // fmt.Println(user.Timestamp)
    fmt.Println("USD Rate: ", user.Rates.USD)
    fmt.Println("EUR Rate: ", user.Rates.EUR)
}
