package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	Log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	Log.SetOutput(os.Stdout)
	Log.Info("Логгер запущен")
}
