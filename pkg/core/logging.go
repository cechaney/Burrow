package core

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//Logger wraps a logrus Logger to provide a common logging config
var Logger = logrus.New()

//ConfigureLogger sets up the log level based on a Viper config
func ConfigureLogger(config *viper.Viper) {

	logFormat := strings.ToUpper(config.GetString("logFormat"))

	if logFormat == "JSON" {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{})
	}

	logLevel := strings.ToUpper(config.GetString("logLevel"))

	switch logLevel {
	case "DEBUG":
		Logger.SetLevel(logrus.DebugLevel)
	case "INFO":
		Logger.SetLevel(logrus.InfoLevel)
	case "WARN":
		Logger.SetLevel(logrus.WarnLevel)
	case "ERROR":
		Logger.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		Logger.SetLevel(logrus.FatalLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
}
