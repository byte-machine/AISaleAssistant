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

func GetHistory(userId string) ([]Message, error) {
	var messages []Message

	rawMessages, err := chat_repos.CheckIfExist(userId)
	if err == nil && len(rawMessages) > 0 {
		messages, err = ParseArrayToArray(rawMessages)

		return messages, err
	}

	return messages, err
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
			Content: `Вы — ассистент, который помогает пользователю с его запросами для покупки в магазине телефонов. Всегда отвечайте понятно и профессионально. ` +
				`Ведите разговоры строго на тему продукта или магазина, при вопросах о твоем создании помни что тебя создала компания Byte-machine. ` +
				//`Прежде чем отвечать о наличии какого либо товара проверь это по бд. ` +
				`Когда клиент просит задает вопрос для ответа на который нужно знать информацию из базы данных, верни в ответ sql запрос, только select ` +
				`В формате: "select * from phones where 'поле' ilike '%значение%' limit 1;" и ничего кроме запроса в текстовом виде. Запрос всегда должен содержать не более 5ти записей. ` +
				`В запросе запрашивай все поля только когда необходимо, и в этих случаях не больше 1й записи, в обычных случаях ограничивайся несколькими полями. ` +
				`Структура таблицы: phone_name(string), brand(string), os(string), inches(float64), resolution(string), battery(int), battery_type(string), ram(int), announcement_date(string), weight(int), storage(int), video_720p(bool), video_1080p(bool), video_4k(bool), video_8k(bool), video_30fps(bool), video_60fps(bool), video_120fps(bool), video_240fps(bool), video_480fps(bool), video_960fps(bool), price_usd(float64) ` +
				`Всегда разделяй название бренда и модель при создании sql запроса в базу. И всегда проверяй на чтобы в запросе не было два одинаковых поля. ` +
				//`Все бренды в базе ` +
				`Если последним сообщением было сообщение от system, не включая самого первого, значит не нужно делать sql запрос, а ответить на последнее сообщение пользователя учитывая данные из этого сообщения. Запрещено отправлять ответы похожие на сообщения от system. `,
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
