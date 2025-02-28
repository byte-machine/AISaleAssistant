package waiting_chat_repos

import (
	"AISale/database"
	. "AISale/database/models"
	"errors"
	"gorm.io/gorm"
	"time"
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

func CheckIfExist(userId string) (WaitingChat, error) {
	db := database.GetDB()
	var chat WaitingChat

	if err := db.Where("chat_user_id = ?", userId).First(&chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return WaitingChat{}, nil
		}
		return WaitingChat{}, err
	}

	return chat, nil
}

func SetIsReminded(userId string, state bool) error {
	db := database.GetDB()

	return db.Model(&WaitingChat{}).Where("chat_user_id = ?", userId).Update("is_reminded", state).Error
}

func FindIfReminded(userId string) (WaitingChat, error) {
	db := database.GetDB()
	var waitingChat WaitingChat

	if err := db.Where("chat_user_id = ? AND is_reminded = ?", userId, true).First(&waitingChat).Error; err != nil {
		return WaitingChat{}, err
	}

	return waitingChat, nil
}

func AddCount(waitingChat *WaitingChat) error {
	db := database.GetDB()

	waitingChat.RemindCount += 1
	waitingChat.Since = time.Now()

	return db.Save(waitingChat).Error
}
