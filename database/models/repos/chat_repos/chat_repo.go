package chat_repos

import (
	"AISale/database"
	. "AISale/database/models"
	"errors"
	"gorm.io/gorm"
)

func CheckIfExist(userId string) ([]string, error) {
	db := database.GetDB()
	var chat Chat

	if err := db.Where("user_id = ?", userId).First(&chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []string{}, nil
		}
		return []string{}, err
	}
	return chat.Messages, nil
}

func SaveChat(userId string, messages []string) error {
	db := database.GetDB()

	var chat Chat
	result := db.First(&chat, "user_id = ?", userId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			chat = Chat{
				UserID:   userId,
				Messages: messages,
			}
			if err := db.Create(&chat).Error; err != nil {
				return errors.New("error creating the record: " + err.Error())
			}
			return nil
		}
		return errors.New("error checking the record: " + result.Error.Error())
	}

	chat.UserID = userId
	chat.Messages = messages
	if err := db.Save(&chat).Error; err != nil {
		return errors.New("error updating the record: " + err.Error())
	}

	return nil
}
