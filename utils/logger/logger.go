package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)

	// TODO: do this from a flag
	Log.SetLevel(logrus.InfoLevel)

	Log.SetFormatter(&logrus.TextFormatter{
		DisableColors:    false,
		DisableTimestamp: true,
		ForceColors:      true,
	})

	Log.WithField("app", "rtt")
}
