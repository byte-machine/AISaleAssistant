package ai_controllers

import (
	"AISale/services/chat"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebhookRequest struct {
	From string `json:"From"`
	Body string `json:"Body"`
}

func SendMessage(c *gin.Context) {
	userId := c.PostForm("user_id")
	userMessage := c.PostForm("user_message")

	response, err := chat.Conservation(c, userId, userMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": response})
}
