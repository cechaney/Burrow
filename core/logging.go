package core

import (
	"strings"

	"github.com/sirupsen/logrus"
)

//Logger wraps a logrus Logger to provide a common logging config
var Logger = initLogging()

func initLogging() *logrus.Logger {

	//TODO: Need to add more appenders
	logger := logrus.New()

	// logger.SetFormatter(&logrus.JSONFormatter{})

	logLevel := Config.GetString("logLevel")

	logLevel = strings.ToUpper(logLevel)

	switch logLevel {
	case "DEBUG":
		logger.SetLevel(logrus.DebugLevel)
	case "INFO":
		logger.SetLevel(logrus.InfoLevel)
	case "WARN":
		logger.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logger.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		logger.SetLevel(logrus.FatalLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	return logger

}
