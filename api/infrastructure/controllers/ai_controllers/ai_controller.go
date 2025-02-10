package ai_controllers

import (
	"AISale/services/chat"
	"AISale/services/twillio"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"log"
	"net/http"
	"strings"
)

type WebhookRequest struct {
	From      string `json:"From"`
	Body      string `json:"Body"`
	NumMedia  string `json:"NumMedia"`
	MediaUrl0 string `json:"MediaUrl0"`
}

func SendMessage(c *gin.Context) {
	userId := c.PostForm("user_id")
	userMessage := c.PostForm("user_message")

	response, err := chat.Conservation(c, userId, userMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": response.Choices[0].Message.Content})
}

func WhatsappWebhook(c *gin.Context) {
	var req WebhookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("error: Invalid request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	log.Printf("💬 Сообщение от %s: %s\n", req.From, req.Body)

	response, err := chat.Conservation(c, req.From, req.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := twillio.SendTwilioMessage(req.From, response.Choices[0].Message.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func SendQuery(c *gin.Context) {
	query := c.PostForm("query")

	queries := strings.Split(query, "|||||")

	var messages []openai.ChatCompletionMessage

	for _, value := range queries {
		chat.AddMessage(&messages, openai.ChatMessageRoleUser, value)
	}

	answer, err := chat.GetAnswer(c, messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": answer.Choices[0].Message.Content})
}
