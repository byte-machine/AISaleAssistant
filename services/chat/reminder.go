package chat

import (
	"AISale/database/models/repos/chat_repos"
	"AISale/database/models/repos/waiting_chat_repos"
	"AISale/services/twillio"
)

func CreateWaitingChat(from string) error {
	chat, err := chat_repos.CheckIfExist(from)
	if err != nil {
		return err
	}

	err = waiting_chat_repos.Create(chat.UserID)
	if err != nil {
		return err
	}

	return nil
}

func Remind(from string) error {
	err := waiting_chat_repos.Delete(from)
	if err != nil {
		return err
	}

	err = twillio.SendTwilioMessage(from, "Вспомните обо мне!")
	if err != nil {
		return err
	}

	return nil
}
