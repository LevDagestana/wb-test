package main

import (
	"os"
	"time"

	"wb/logger"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:9092"}
	topic := "order_topic"

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	producer, err := sarama.NewSyncProducer(brokers, config)
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
