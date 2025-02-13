package log

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() {
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

func info(message string, fields map[string]interface{}) {
	log.WithFields(fields).Info(message)

}

func error(message string, fields map[string]interface{}) {
	log.WithFields(fields).Error(message)

}
