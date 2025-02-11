package chat

import (
	"AISale/database/models/repos/chat_repos"
	"errors"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/sashabaranov/go-openai"
	"log"
	"strings"
)

type Message struct {
	Role    string
	Content string
}
type Chat struct {
	UserID   string
	Messages []Message
}

func GetHistory(userId string) ([]Message, error) {
	var messages []Message

	rawMessages, err := chat_repos.CheckIfExist(userId)
	if err == nil && len(rawMessages) > 0 {
		messages, err = ParseArrayToArray(rawMessages)

		return messages, err
	}

	return messages, err
}

func GetAllChats() ([]Chat, error) {
	var parsedChats []Chat

	chats, err := chat_repos.GetAllChats()
	if err == nil && len(chats) > 0 {
		for _, chat := range chats {
			messages, err := ParseArrayToArray(chat.Messages)
			parsedChats = append(parsedChats, Chat{UserID: chat.UserID, Messages: messages})

			return parsedChats, err
		}

		return parsedChats, err
	}

	return parsedChats, err
}

func GetMessages(userId string) ([]openai.ChatCompletionMessage, error) {
	var messages []openai.ChatCompletionMessage

	rawMessages, err := chat_repos.CheckIfExist(userId)
	if err != nil {
		return messages, err
	} else if len(rawMessages) == 0 {
		messages = StartMessages()
	} else {
		messages, err = ParseArrayToMessages(rawMessages)

		return messages, err
	}

	return messages, nil
}

func StartMessages() []openai.ChatCompletionMessage {
	log.Printf("Принял системный промпт")
	messages := []openai.ChatCompletionMessage{
		{
			Role: openai.ChatMessageRoleSystem,
			Content: `Ты профессиональный менеджер по продажам, предлагающий услуги обучения в компании Атамекен. ` +
				`В начале диалога если пользователь не задал прямых вопросов по курсам, начни задавать наводящие вопросы. ` +
				`Не допускай пустых разговоров со своей стороны, всегда предлагай услуги. ` +
				`Обязательно нужно чтобы пользователь ознакомился с услугами которые мы предоставляем а затем узнал цену обучения. ` +
				`Собирай информацию об количестве человек для курса и подводи итог по цене. ` +
				`Ведите разговоры строго на тему услуг, при вопросах о твоем создании помни что тебя создала компания Byte-machine. `,
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
