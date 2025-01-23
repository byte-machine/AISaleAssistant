package chat_controllers

import (
	"AISale/services/chat"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetChatHistory(c *gin.Context) {
	userId := c.PostForm("user_id")

	messages, err := chat.GetHistory(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": messages})
}
