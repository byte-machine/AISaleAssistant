package chat

import (
	"AISale/database/models/repos/chat_repos"
	"errors"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/sashabaranov/go-openai"
	"strings"
)

func GetMessages(userId string) ([]openai.ChatCompletionMessage, error) {
	var messages []openai.ChatCompletionMessage

	rawMessages, err := chat_repos.CheckIfExist(userId)
	if err != nil {
		return messages, err
	} else if len(rawMessages) == 0 {
		messages = StartMessages()
	} else {
		messages, err = ParseArrayToMessages(rawMessages)
		if err != nil {
			return messages, err
		}
	}

	return messages, nil
}

func StartMessages() []openai.ChatCompletionMessage {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Вы — ассистент, который помогает пользователю с его запросами для покупки в магазине телефонов. Всегда отвечайте понятно и профессионально.",
		},
	}

	return messages
}

func AddMessage(messages *[]openai.ChatCompletionMessage, role string, message string) {
	*messages = append(*messages, openai.ChatCompletionMessage{Role: role, Content: message})
}

func SaveMessages(userId string, messages []openai.ChatCompletionMessage) error {
	arrayMessages := SerializeMessagesToArray(messages)

	err := chat_repos.SaveChat(userId, arrayMessages)
	if err != nil {
		return err
	}

	return nil
}

func SerializeMessagesToArray(messages []openai.ChatCompletionMessage) []string {
	var arrayMessages []string

	for _, value := range messages {
		role := value.Role
		content := value.Content

		arrayMessages = append(arrayMessages, role+"||"+content)
	}

	return arrayMessages
}

func ParseArrayToMessages(arrayMessages []string) ([]openai.ChatCompletionMessage, error) {
	var messages []openai.ChatCompletionMessage

	for _, value := range arrayMessages {
		parts := strings.SplitN(value, "||", 2)

		set := mapset.NewSet("system", "assistant", "user")
		if len(parts) == 2 && set.Contains(parts[0]) {
			messages = append(messages, openai.ChatCompletionMessage{
				Role:    parts[0],
				Content: parts[1],
			})
		} else {
			return []openai.ChatCompletionMessage{}, errors.New("messages have wrong structure, incorrect user type")
		}
	}

	return messages, nil
}
