package rest

import (
	"AISale/api/infrastructure/controllers/ai_controllers"
	"AISale/config"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

var client *openai.Client

func GetAIClient() *openai.Client {
	return client
}

func AIRoutes(router *gin.Engine, settings config.Settings) {
	client = openai.NewClient(settings.OpenaiApiKey)

	productGroup := router.Group("chat")
	{
		productGroup.POST("/user_message", ai_controllers.GetMessage)
	}
}
