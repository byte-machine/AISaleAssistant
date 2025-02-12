package database

import (
	"AISale/config"
	. "AISale/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func Connect(settings config.Settings) {
	host := settings.DbHost
	user := settings.DbUser
	password := settings.DbPassword
	dbname := settings.DbName
	port := settings.DbPort
	sslmode := settings.Ssl

	var err error
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Автоматическое создание таблиц
	err = db.AutoMigrate(&Chat{})
	if err != nil {
		log.Fatalf("Не удалось создать таблицы: %v", err)
	}

	log.Println("Таблицы успешно созданы или уже существуют")
}

func Disconnect() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatal("Failed to close the database connection:", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
