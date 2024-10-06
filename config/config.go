package config

import (
	"wb/logger"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Kafka    KafkaConfig    `mapstructure:"kafka"`
	Http     HttpConfig     `mapstructure:"http"`
}

type DatabaseConfig struct {
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	DBName     string `mapstructure:"dbname"`
	SSLMode    string `mapstructure:"sslmode"`
	DriverName string `mapstructure:"driver_name"`
}

type KafkaConfig struct {
	Brokers []string `mapstructure:"brokers"`
	Topic   string   `mapstructure:"topic"`
	GroupID string   `mapstructure:"group_id"`
}
type HttpConfig struct {
	Port string `mapstructure:"port"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		logger.Log.WithError(err).Fatalf("Ошибка при чтении конфигурации:")
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		logger.Log.WithError(err).Fatalf("Ошибка при декодировании конфигурации: ")
		return
	}

	return config, nil
}
