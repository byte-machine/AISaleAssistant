package logger

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
}

func Info(message string, fields map[string]interface{}) {
	log.WithFields(fields).Info(message)
}

func Error(message string, fields map[string]interface{}) {
	log.WithFields(fields).Error(message)
}
