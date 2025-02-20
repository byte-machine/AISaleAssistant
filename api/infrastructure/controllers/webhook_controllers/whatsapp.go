package webhook_controllers

import (
	"AISale/services/chat"
	"AISale/services/twillio"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func WhatsappAnswer(c *gin.Context) {
	from := c.PostForm("From")
	body := c.PostForm("Body")

	log.Printf("💬 Сообщение от %s: %s\n", from, body)

	response, err := chat.Conservation(c, from, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = twillio.SendTwilioMessage(from, response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func WhatsappReminderStart(c *gin.Context) {
	from := c.PostForm("From")
	status := c.PostForm("SmsStatus")

	if status == "delivered" {
		err := chat.CreateWaitingChat(from)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.Status(http.StatusOK)
}
