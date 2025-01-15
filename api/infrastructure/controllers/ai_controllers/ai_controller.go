package ai_controllers

import (
	"AISale/api/infrastructure/rest"
	"AISale/database/models/repos/chat_repos"
	"AISale/services/chat"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

func GetMessage(c *gin.Context) {
	userId := c.PostForm("user_id")
	client := rest.GetAIClient()

	var messages []openai.ChatCompletionMessage

	rawMessages, err := chat_repos.CheckIfExist(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if len(rawMessages) == 0 {
		messages = chat.StartMessages()
	} else {
		messages, err = chat.ParseArrayToMessages(rawMessages)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	response, err := client.CreateChatCompletion(c, openai.ChatCompletionRequest{
		Model:    openai.GPT4,
		Messages: messages,
	})

	c.JSON(http.StatusCreated, gin.H{"status": status, "book": book})
	return
}
