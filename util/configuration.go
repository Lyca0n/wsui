package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func GetLogLevel() log.Level {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	switch logLevel {
	case "debug":
		log.SetReportCaller(true)
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	case "panic":
		return log.PanicLevel
	default:
		return log.InfoLevel
	}
}
