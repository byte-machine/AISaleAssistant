package ai_controllers

import (
	"AISale/logger"
	"AISale/services/chat"
	"AISale/services/twillio"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookRequest struct {
	From string `json:"From"`
	Body string `json:"Body"`
}

func SendMessage(c *gin.Context) {
	userId := c.PostForm("user_id")
	userMessage := c.PostForm("user_message")

	logger.Info("📩 Новое сообщение от пользователя", map[string]interface{}{
		"userId":      userId,
		"userMessage": userMessage,
	})

	response, err := chat.Conservation(c, userId, userMessage)
	if err != nil {
		logger.Error("Error processing message", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Отправка сообщения пользователю", map[string]interface{}{
		"userId":   userId,
		"response": response,
	})

	c.JSON(http.StatusOK, gin.H{"answer": response})
}

func WhatsappWebhook(c *gin.Context) {
	from := c.PostForm("From")
	body := c.PostForm("Body")

	log.Printf("💬 Сообщение от %s: %s\n", from, body)

	response, err := chat.Conservation(c, from, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := twillio.SendTwilioMessage(from, response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
