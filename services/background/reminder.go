package background

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
			if time.Since(waitingChat.Since) >= config.WaitingTime {
				fmt.Println("Прошел 1 час, напоминаем пользователю!")

				err = chat.Remind(waitingChat.ChatUserID)
				if err != nil {
					continue
				}
			}
		}

		time.Sleep(1 * time.Minute)
	}
}
