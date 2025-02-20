package rest

import (
	"AISale/api/infrastructure/controllers/webhook_controllers"
	"AISale/config"
	"github.com/gin-gonic/gin"
)

func WebHookRoutes(router *gin.Engine, settings config.Settings) {
	productGroup := router.Group("webhook")
	{
		productGroup.POST("/whatsapp_answer", webhook_controllers.WhatsappAnswer)
		productGroup.POST("/whatsapp_delivered", webhook_controllers.WhatsappReminderStart)
	}
}
