package config

import "github.com/sashabaranov/go-openai"

var client *openai.Client

func GetAIClient() *openai.Client {
	return client
}

func InitClient(openaiApiKey string) {
	client = openai.NewClient(openaiApiKey)
}
