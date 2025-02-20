package chat

import (
	"AISale/config"
	. "AISale/database/models"
	"AISale/database/models/repos/chat_repos"
	"errors"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/sashabaranov/go-openai"
	"log"
	"strings"
	"time"
)

//type Message struct {
//	Role    string
//	Content string
//}
//type Chat struct {
//	UserID   string
//	Messages []Message
//}

func GetHistory(userId string) ([]Message, error) {
	chat, err := chat_repos.CheckIfExist(userId)

	return chat.Messages, err
}

func GetAllChats() ([]string, error) {
	var parsedChats []string

	chats, err := chat_repos.GetAll()
	if err == nil && len(chats) > 0 {
		for _, chat := range chats {

			parsedChats = append(parsedChats, chat.UserID)
		}

		return parsedChats, err
	}

	return parsedChats, err
}

func GetMessages(userId string) ([]openai.ChatCompletionMessage, error) {
	var messages []openai.ChatCompletionMessage

	chat, err := chat_repos.CheckIfExist(userId)
	rawMessages := chat.Messages

	if err != nil {
		return messages, err
	} else if len(rawMessages) == 0 {
		messages = StartMessages()
	} else {
		messages, err = ConvertToOpenaiMessage(rawMessages)

		CheckSystemMessages(&messages)

		return messages, err
	}

	return messages, nil
}

func StartMessages() []openai.ChatCompletionMessage {
	log.Printf("Принял системный промпт")

	return config.Messages
}

func AddMessage(messages *[]openai.ChatCompletionMessage, role string, message string) {
	*messages = append(*messages, openai.ChatCompletionMessage{Role: role, Content: message})
}

func SaveMessages(userId string, messages []openai.ChatCompletionMessage) error {
	arrayMessages := ConvertToMessage(messages)

	err := chat_repos.Save(userId, arrayMessages)
	if err != nil {
		return err
	}

	return nil
}

func ConvertToMessage(messages []openai.ChatCompletionMessage) []Message {
	var messagesArray []Message

	for _, message := range messages {
		messagesArray = append(messagesArray, Message{
			Role:    message.Role,
			Content: message.Content,
			Time:    time.Now(),
		})
	}

	return messagesArray
}

func ConvertToOpenaiMessage(arrayMessages []Message) ([]openai.ChatCompletionMessage, error) {
	var messages []openai.ChatCompletionMessage

	for _, message := range arrayMessages {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	return messages, nil
}

func ParseArrayToArray(arrayMessages []string) ([]Message, error) {
	var messages []Message

	for _, value := range arrayMessages {
		parts := strings.SplitN(value, "||", 2)

		set := mapset.NewSet("system", "assistant", "user")
		if parts[0] == "system" {
			continue
		} else if len(parts) == 2 && set.Contains(parts[0]) {
			messages = append(messages, Message{Role: parts[0], Content: parts[1]})
		} else {
			return []Message{}, errors.New("messages have wrong structure, incorrect user type")
		}
	}

	return messages, nil
}

func CheckSystemMessages(messages *[]openai.ChatCompletionMessage) {
	var systemMessages, otherMessages []openai.ChatCompletionMessage

	for _, message := range *messages {
		if message.Role == "system" {
			systemMessages = append(systemMessages, message)
		} else {
			otherMessages = append(otherMessages, message)
		}
	}

	var isUpdated = true
	if len(config.Messages) != len(systemMessages) {
		isUpdated = false
	} else {
		for num, message := range config.Messages {
			if systemMessages[num].Content != message.Content {
				isUpdated = false
				break
			}
		}
	}

	if !isUpdated {
		*messages = config.Messages
		*messages = append(*messages, otherMessages...)
	}
}
