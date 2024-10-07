package kafka

import (
	"context"
	"encoding/json"

	"wb/models"
	"wb/repository"

	"wb/config"

	"wb/logger"

	"github.com/IBM/sarama"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type consumerGroupHandler struct{}

func InitKafka(cfg *config.Config) {

	saramaCfg := sarama.NewConfig()
	saramaCfg.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	saramaCfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	brokers := []string{cfg.KafkaBroker}
	group := cfg.KafkaGroupID
	topic := cfg.KafkaTopic

	logger.Log.WithFields(logrus.Fields{
		"brokers": brokers,
		"group":   group,
	}).Info("Создаем группу потребителей...")

	consumerGroup, err := sarama.NewConsumerGroup(brokers, group, saramaCfg)
	if err != nil {
		logger.Log.WithError(err).Fatal("Ошибка при создании группы потребителей:")
	}

	logger.Log.Info("Группа потребителей успешно создана")

	handler := &consumerGroupHandler{}

	go func() {
		for {
			if err := consumerGroup.Consume(context.Background(), []string{topic}, handler); err != nil {
				logger.Log.WithError(err).Fatal("Ошибка при подписке на тему:")
			}
		}
	}()

}

func (handler *consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (handler *consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (handler *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		logger.Log.WithFields(logrus.Fields{
			"topic":   claim.Topic(),
			"offset":  message.Offset,
			"message": string(message.Value),
		}).Info("Получено сообщение из Kafka")

		var order models.Order
		err := json.Unmarshal(message.Value, &order)
		if err != nil {
			logger.Log.WithError(err).Error("Ошибка при десериализации сообщения:")
			continue
		}
		validate := validator.New()
		err = validate.Struct(order)
		if err != nil {
			logger.Log.WithError(err).Error("Ошибка при валидации заказа")
			continue
		}
		repository.InsertOrder(order)

		session.MarkMessage(message, "")
	}
	return nil
}
