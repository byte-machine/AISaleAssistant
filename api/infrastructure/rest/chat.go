package rest

import (
	"AISale/api/infrastructure/controllers/ai_controllers"
	"AISale/api/infrastructure/controllers/chat_controllers"
	"AISale/config"
	"github.com/gin-gonic/gin"
)

func ChatRoutes(router *gin.Engine, settings config.Settings) {
	productGroup := router.Group("chat")
	{
		productGroup.POST("/send_message", ai_controllers.SendMessage)
		productGroup.POST("/whatsapp_webhook", ai_controllers.WhatsappWebhook)
		productGroup.GET("/get_chats", chat_controllers.GetChats)
		productGroup.POST("/get_chat", chat_controllers.GetChatHistory)
	}
}
