package rest

import (
	"AISale/api/infrastructure/controllers/ai_controllers"
	"AISale/config"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine, settings config.Settings) {
	productGroup := router.Group("chat")
	{
		productGroup.POST("/send_query", ai_controllers.SendQuery)
		// productGroup.POST("/get", ai_controllers.)
	}
}
