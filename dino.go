package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/dino/dynowebportal"
)

type configuration struct {
	WebServer string `json:"webServer"`
}

func main() {
	file, err := os.Open("config.json")

	if err != nil {
		log.Fatal(err)
	}

	config := new(configuration)

	json.NewDecoder(file).Decode(config)

	log.Println("starting web-server on PORT: ", config.WebServer)

	dynowebportal.RunWebPortal(config.WebServer)
}
