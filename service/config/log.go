package config

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// SetupLogger настраивает и возвращает логгер на основе уровня логирования и формата.
func SetupLogger(level, format string) *logrus.Logger {
	logger := logrus.New()

	// Настраиваем уровень логирования.
	lvl, err := logrus.ParseLevel(strings.ToLower(level))
	if err != nil {
		logger.Warnf("invalid log level '%s', defaulting to 'info'", level)
		lvl = logrus.InfoLevel
	}
	logger.SetLevel(lvl)

	// Настраиваем формат логирования.
	switch strings.ToLower(format) {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		fallthrough
	default:
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	return logger
}
