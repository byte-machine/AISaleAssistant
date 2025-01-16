package rest

import (
	"AISale/api/infrastructure/controllers/ai_controllers"
	"AISale/config"
	"github.com/gin-gonic/gin"
)

func AIRoutes(router *gin.Engine, settings config.Settings) {
	productGroup := router.Group("chat")
	{
		productGroup.POST("/send_message", ai_controllers.SendMessage)
	}
}
