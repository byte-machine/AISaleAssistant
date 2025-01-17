package phone_repos

import (
	"AISale/database"
	. "AISale/database/models"
	"encoding/json"
	"errors"
)

func Create(phone Phone) error {
	db := database.GetDB()

	if err := db.Create(&phone).Error; err != nil {
		return errors.New("невозможно создать новый объект")
	}

	return nil
}

func RawSelect(query string) (string, error) {
	db := database.GetDB()

	var results []map[string]interface{}

	// Выполнение запроса
	result := db.Raw(query).Scan(&results)
	if result.Error != nil {
		return "", result.Error
	}

	// Преобразуем результат в JSON, чтобы оставить только заполненные поля
	jsonData, err := json.Marshal(results)
	if err != nil {
		return "", err
	}

	// Возвращаем результат в текстовом формате
	return string(jsonData), nil
}
