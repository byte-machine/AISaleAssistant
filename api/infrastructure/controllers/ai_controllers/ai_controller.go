package ai_controllers

import (
	"AISale/database/models/repos/chat_repos"
	"AISale/services/chat"
	"AISale/services/twillio"
	"github.com/gin-gonic/gin"
	"log"
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

	c.JSON(http.StatusOK, gin.H{"answer": response.Choices[0].Message.Content})
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

	if response.Choices[0].Message.Content != "ending" {
		if err := twillio.SendTwilioMessage(from, response.Choices[0].Message.Content); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := chat_repos.SetClientStatusTrue(from); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.Status(http.StatusOK)
}
