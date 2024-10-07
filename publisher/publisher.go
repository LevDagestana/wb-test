package main

import (
	"os"
	"time"

	"wb/config"
	"wb/logger"

	"github.com/IBM/sarama"
)

func main() {
	cfg, err := config.LoadConfig("..")
	if err != nil {
		logger.Log.WithError(err).Fatal("Ошибка загрузки конфигурации:")
	}
	brokers := []string{cfg.KafkaBroker}
	topic := cfg.KafkaTopic
	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Return.Successes = true
	saramaCfg.Producer.Timeout = 5 * time.Second

	producer, err := sarama.NewSyncProducer(brokers, saramaCfg)
	if err != nil {

		logger.Log.WithError(err).Fatal("Ошибка при создании Kafka producer:")
	}
	defer producer.Close()

	data, err := os.ReadFile("order.json")
	if err != nil {
		logger.Log.WithError(err).Fatalf("Failed to read JSON file:")
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	_, _, err = producer.SendMessage(message)
	if err != nil {
		logger.Log.WithError(err).Fatal("Ошибка при отправке сообщения:")
	}

	logger.Log.Info("Сообщение успешно отправлено")

}
