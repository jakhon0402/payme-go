package logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Config struct {
}

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339, // Custom timestamp format
		FullTimestamp:   true,
	})
	logger.SetOutput(os.Stdout)
	return logger
}
