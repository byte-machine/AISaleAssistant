package chat

import (
	"AISale/config"
	"AISale/services/external_api"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"log"
	"strings"
)

func GetAnswer(c *gin.Context, messages []openai.ChatCompletionMessage, consType config.ConservationType) (openai.ChatCompletionResponse, error) {
	client := config.GetAIClient()

	var model string
	if consType == config.Bytemachine {
		model = "gpt-3.5-turbo"
	} else {
		model = "ft:gpt-3.5-turbo-0125:personal::B07BtIZ4"
	}

	response, err := client.CreateChatCompletion(c, openai.ChatCompletionRequest{
		Model:    model,
		Messages: messages,
	})
	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}

	return response, nil
}

func Conservation(c *gin.Context, userId string, userMessage string, consType config.ConservationType) (string, error) {
	messages, err := GetMessages(userId, consType)
	if err != nil {
		return "", err
	}

	if consType == config.Remind {
		AddMessage(&messages, "system", userMessage)
	} else {
		AddMessage(&messages, "user", userMessage)
	}

	response, err := GetAnswer(c, messages, consType)
	if err != nil {
		return "", err
	}

	jsonText := GetJSONFromText(response.Choices[0].Message.Content)
	log.Printf("Наличие json в смтроке %s\n", jsonText)
	if len(jsonText) != 0 {
		response.Choices[0].Message.Content = strings.TrimSpace(strings.Replace(response.Choices[0].Message.Content, jsonText, "", 1))

		log.Printf("Данные для записи пользователя %s: %s\n", userId, jsonText)

		external_api.SendToAli(jsonText)
	}

	AddMessage(&messages, "assistant", response.Choices[0].Message.Content)

	err = SaveMessages(userId, messages)
	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
