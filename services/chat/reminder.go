package chat

import (
	"AISale/config"
	"AISale/database/models/repos/waiting_chat_repos"
	"AISale/services/twillio"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
)

func CreateWaitingChat(to string) error {
	err := waiting_chat_repos.Create(to)
	if err != nil {
		return err
	}

	return nil
}

func Remind(userId string) error {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	message, err := Conservation(c, userId, "пользователь долго не отвечает, возобнови беседу сообщением такого типа: '\nСогласитесь, легче заплатить 120000 (на 10 человек) и обезопасить себя от проверок на целый год, чем рисковать попасть под штраф в 750 000 или еще хуже, понести уголовную ответственность при несчастном случае'", config.Remind)
	if err != nil {
		return err
	}

	err = twillio.SendTwilioMessage(userId, message)
	if err != nil {
		return err
	}

	err = waiting_chat_repos.SetIsRemindedTrue(userId)
	if err != nil {
		return err
	}

	return nil
}
