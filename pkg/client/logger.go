package client

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewFileLogger creates a log file and init logger
func NewFileLogger(path string) *logrus.Logger {
	logger := logrus.New()

	// default log to os.Stderr
	if path == "" {
		return logger
	}

	if file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666); err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}

	// use json formatter
	logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	return logger
}
