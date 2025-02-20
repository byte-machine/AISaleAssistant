package waiting_chat_repos

import (
	"AISale/database"
	. "AISale/database/models"
	"errors"
)

func Create(userId string) error {
	db := database.GetDB()

	waitingChat := WaitingChat{
		ChatUserID: userId,
	}
	if err := db.Create(&waitingChat).Error; err != nil {
		return errors.New("error creating the record: " + err.Error())
	}
	return nil
}

func Delete(userId string) error {
	db := database.GetDB()

	waitingChat := WaitingChat{
		ChatUserID: userId,
	}

	if err := db.Where(&waitingChat).Delete(&waitingChat).Error; err != nil {
		return errors.New("error deleting the record: " + err.Error())
	}

	return nil
}

func GetAll() ([]WaitingChat, error) {
	db := database.GetDB()

	var waitingChats []WaitingChat

	if err := db.Find(&waitingChats).Error; err != nil {
		return nil, errors.New("error getting the records: " + err.Error())
	}

	return waitingChats, nil
}
