package webhook_controllers

import (
	"AISale/database/models"
	"AISale/database/models/repos/waiting_chat_repos"
	"AISale/services/chat"
	"AISale/services/twillio"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func WhatsappAnswer(c *gin.Context) {
	from := c.PostForm("From")
	body := c.PostForm("Body")

	log.Printf("💬 Сообщение от %s: %s\n", from, body)

	err := waiting_chat_repos.Delete(from)
	if err != nil {
		log.Printf("waiting chat deleting error: %s\n", err.Error())
	}

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
	to := c.PostForm("To")
	status := c.PostForm("SmsStatus")

	fmt.Printf("Добавление чата пользователя %s!\n", to)
	fmt.Printf("Статус %s!\n", status)

	if status == "delivered" {
		exist, err := waiting_chat_repos.CheckIfExist(to)
		if err != nil {
			return
		}
		if exist == (models.WaitingChat{}) {
			err = chat.CreateWaitingChat(to)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	c.Status(http.StatusOK)
}
