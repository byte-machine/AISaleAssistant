package ai_controllers

import (
	"AISale/database/models/repos/phone_repos"
	"AISale/services/chat"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"log"
	"net/http"
	"strings"
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

	if strings.Contains(strings.ToLower(response.Choices[0].Message.Content), "select") {
		log.Println("sql query: " + response.Choices[0].Message.Content)

		query := response.Choices[0].Message.Content
		phones, err := phone_repos.RawSelect(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Println("query answer: " + phones)

		chat.AddMessage(&messages, "system", "Ответь на последний запрос пользователя. \n\nответ на sql запрос '"+query+"' - "+phones)

		response, err = chat.GetAnswer(c, messages)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	chat.AddMessage(&messages, "assistant", response.Choices[0].Message.Content)

	err = chat.SaveMessages(userId, messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": response.Choices[0].Message.Content})
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
