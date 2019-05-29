package core

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//Logger wraps a logrus Logger to provide a common logging config
var Logger = logrus.New()

func init() {

}

//ConfigureLogger sets up the log level based on a Viper config
func ConfigureLogger(config *viper.Viper) {

	logFormat := config.GetString("logFormat")

	logFormat = strings.ToUpper(logFormat)

	if logFormat == "JSON" {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{})
	}

	logLevel := config.GetString("logLevel")

	logLevel = strings.ToUpper(logLevel)

	if len(logLevel) == 0 {
		logLevel = "INFO"
	}

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
