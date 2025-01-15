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
