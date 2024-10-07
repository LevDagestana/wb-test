package config

import (
	"os"
	"path/filepath"
	"wb/logger"

	"github.com/joho/godotenv"
)

// Config структура для хранения конфигураций
type Config struct {
	DbDriverName string
	DbScheme     string
	KafkaBroker  string
	KafkaTopic   string
	KafkaGroupID string
	Port         string
}

func LoadConfig(path string) (*Config, error) {
	dir, err := os.Getwd()
	if err != nil {
		logger.Log.WithError(err).Fatal("")
	}
	err = godotenv.Load(filepath.Join(dir, path, ".env"))

	if err != nil {
		logger.Log.WithError(err).Error("Ошибка загрузки .env файла:")
	}

	cfg := &Config{
		DbDriverName: os.Getenv("DB_DRIVER_NAME"),
		DbScheme:     os.Getenv("DB_SCHEME"),
		KafkaBroker:  os.Getenv("KAFKA_BROKER"),
		KafkaTopic:   os.Getenv("KAFKA_TOPIC"),
		KafkaGroupID: os.Getenv("KAFKA_GROUP_ID"),
		Port:         os.Getenv("PORT"),
	}

	return cfg, nil
}
