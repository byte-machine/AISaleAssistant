package webhook_controllers

import (
	"AISale/config"
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

	//err := waiting_chat_repos.Delete(from)
	//if err != nil {
	//	log.Printf("waiting chat deleting error: %s\n", err.Error())
	//}
	if reminded, err := waiting_chat_repos.FindIfReminded(from); err != nil {
		log.Printf("error find waiting chat: %s\n", err.Error())
	} else {
		if err = waiting_chat_repos.AddCount(&reminded); err != nil {
			log.Printf("error add count to waiting chat: %s\n", err.Error())
		}
		if err = waiting_chat_repos.SetIsReminded(reminded.ChatUserID, false); err != nil {
			log.Printf("error set is_reminded of waiting chat: %s\n", err.Error())
		}
	}

	response, err := chat.Conservation(c, from, body, config.Twilio)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("conservation error: %s\n", err.Error())
		return
	}

	if err = twillio.SendTwilioMessage(from, response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("send to twillio error: %s\n", err.Error())
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

		if exist, err := waiting_chat_repos.CheckIfExist(to); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else if exist == (models.WaitingChat{}) {
			err = chat.CreateWaitingChat(to)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	c.Status(http.StatusOK)
}
