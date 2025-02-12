package api

import (
	"AISale/api/infrastructure/rest"
	"AISale/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func RouterStart(settings config.Settings) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		MaxAge:       12 * 60 * 60,
	}))

	rest.ChatRoutes(r, settings)

	err := r.Run(":8081")
	if err != nil {
		log.Fatal("Router start error", err)
	}
}
