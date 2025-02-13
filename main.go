package main

import (
	"AISale/api"
	"AISale/config"
	"AISale/database"
	"log"

	"gorm.io/gorm/logger"
)

func main() {
	logger.InitLogger()

	settings, err := config.LoadENV()
	if err != nil {
		logger.Error("Failed to load environment variables", map[string]interface{}{
			"error": err.Error(),
		})

		log.Fatal(err)
	}
	config.InitClient(settings.OpenaiApiKey)

	database.Connect(settings)
	defer database.Disconnect()

	logger.Info("Starting api server", nil)

	api.RouterStart(settings)
}
