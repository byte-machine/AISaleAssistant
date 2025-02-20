package chat

import (
	"AISale/database/models/repos/waiting_chat_repos"
	"AISale/services/twillio"
)

func CreateWaitingChat(to string) error {
	err := waiting_chat_repos.Create(to)
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
