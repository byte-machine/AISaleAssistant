package chat

import (
	"AISale/config"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

func GetAnswer(c *gin.Context, messages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	client := config.GetAIClient()

	response, err := client.CreateChatCompletion(c, openai.ChatCompletionRequest{
		Model:    "ft:gpt-3.5-turbo-0125:personal::AzJRcq4v",
		Messages: messages,
	})
	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}

	return response, nil
}

func Conservation(c *gin.Context, userId string, userMessage string) (openai.ChatCompletionResponse, error) {
	messages, err := GetMessages(userId)
	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}

	AddMessage(&messages, "user", userMessage)

	response, err := GetAnswer(c, messages)
	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}

	AddMessage(&messages, "assistant", response.Choices[0].Message.Content)

	err = SaveMessages(userId, messages)
	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}

	return response, nil
}
