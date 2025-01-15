package main

import (
	"AISale/api"
	"AISale/config"
	"AISale/database"
	"log"
)

func main() {
	settings, err := config.LoadENV()
	if err != nil {
		log.Fatal(err)
	}

	database.Connect(settings)
	defer database.Disconnect()

	api.RouterStart(settings)
}
