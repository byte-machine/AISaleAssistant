package chat

import (
	"AISale/config"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

func GetAnswer(c *gin.Context, messages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	client := config.GetAIClient()

	response, err := client.CreateChatCompletion(c, openai.ChatCompletionRequest{
		Model:    openai.GPT4o,
		Messages: messages,
	})
	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}

	return response, nil
}
