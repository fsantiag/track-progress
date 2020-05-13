package configuration

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// Logger is a singleton to information logs on the project
	Logger logrus.Logger
)

func init() {
	Logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.InfoLevel)
}
