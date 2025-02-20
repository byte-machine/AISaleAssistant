package chat

import (
	"AISale/config"
	"AISale/database/models/repos/chat_repos"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"strings"
)

func GetAnswer(c *gin.Context, messages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	client := config.GetAIClient()

	response, err := client.CreateChatCompletion(c, openai.ChatCompletionRequest{
		Model:    "ft:gpt-3.5-turbo-0125:personal::B07BtIZ4",
		Messages: messages,
	})
	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}

	return response, nil
}

func Conservation(c *gin.Context, userId string, userMessage string) (string, error) {
	messages, err := GetMessages(userId)
	if err != nil {
		return "", err
	}

	AddMessage(&messages, "user", userMessage)

	response, err := GetAnswer(c, messages)
	if err != nil {
		return "", err
	}

	if strings.Contains(response.Choices[0].Message.Content, "ending") {
		if err := chat_repos.SetClientStatusTrue(userId); err != nil {
			return "", err
		}
		//response.Choices[0].Message.Content = "Отлично, мы позвоним вам в ближайшее время для совершения оплаты услуг."

		response.Choices[0].Message.Content = strings.ReplaceAll(response.Choices[0].Message.Content, "ending", "")
		response.Choices[0].Message.Content = strings.ReplaceAll(response.Choices[0].Message.Content, "|", "")
		if len(response.Choices[0].Message.Content) <= 5 {
			response.Choices[0].Message.Content = "Отлично, мы позвоним вам в ближайшее время для совершения оплаты услуг."
		}
	}

	AddMessage(&messages, "assistant", response.Choices[0].Message.Content)

	err = SaveMessages(userId, messages)
	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
