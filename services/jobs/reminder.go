package jobs

import (
	"AISale/config"
	"AISale/database/models/repos/waiting_chat_repos"
	"AISale/services/chat"
	"fmt"
	"time"
)

func CheckWaitingChats() {
	for {
		waitingChats, err := waiting_chat_repos.GetAll()
		if err != nil {
			continue
		}

		for _, waitingChat := range waitingChats {
			if !waitingChat.IsReminded && waitingChat.RemindCount < config.MaxRemindCount && time.Since(waitingChat.Since) >= config.WaitingTime {
				fmt.Printf("Прошел 1 час, напоминаем пользователю %s!\n", waitingChat.ChatUserID)

				err = chat.Remind(waitingChat.ChatUserID)
				if err != nil {
					fmt.Printf("Произошла ошибка во время напоминания: %s", waitingChat.ChatUserID)
					continue
				}
			}
		}

		time.Sleep(1 * time.Minute)
	}
}
