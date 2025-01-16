package ai_controllers

import (
	"AISale/services/chat"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendMessage(c *gin.Context) {
	userId := c.PostForm("user_id")
	userMessage := c.PostForm("user_message")

	messages, err := chat.GetMessages(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chat.AddMessage(&messages, "user", userMessage)

	response, err := chat.GetAnswer(c, messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	chat.AddMessage(&messages, "assistant", response.Choices[0].Message.Content)

	err = chat.SaveMessages(userId, messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": response.Choices[0].Message.Content})
	return
}
