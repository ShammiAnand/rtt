package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func logLevelFromString(logLevel string) logrus.Level {
	switch logLevel {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "ERROR":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}

func InitLogger(logLevel string) {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)

	Log.SetLevel(logLevelFromString(logLevel))

	Log.SetFormatter(&logrus.TextFormatter{
		DisableColors:    false,
		DisableTimestamp: true,
		ForceColors:      true,
	})

	Log.WithField("app", "rtt")
}
