package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Settings struct {
	DbHost       string
	DbUser       string
	DbPassword   string
	DbName       string
	DbPort       string
	Ssl          string
	OpenaiApiKey string
}

func LoadENV() (Settings, error) {
	err := godotenv.Load()
	if err != nil {
		return Settings{}, errors.New("error loading .env file: " + err.Error())
	}

	settings := Settings{
		DbHost:       os.Getenv("DB_HOST"),
		DbUser:       os.Getenv("DB_USER"),
		DbPassword:   os.Getenv("DB_PASSWORD"),
		DbName:       os.Getenv("DB_NAME"),
		DbPort:       os.Getenv("DB_PORT"),
		Ssl:          os.Getenv("DB_SSL"),
		OpenaiApiKey: os.Getenv("OPENAI_API_KEY"),
	}

	return settings, nil
}
