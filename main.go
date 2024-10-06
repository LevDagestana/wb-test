package main

import (
	"net/http"

	"wb/handlers"
	"wb/kafka"
	"wb/repository"

	"wb/db"

	"wb/config"

	"wb/logger"

	_ "github.com/lib/pq"
)

func main() {
	logger.InitLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Log.WithError(err).Fatal("Ошибка загрузки конфигурации:")
	}
	db.InitDb(cfg.Database)
	kafka.InitKafka(cfg.Kafka)
	repository.LoadCache()
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/order", handlers.GetOrderByIdHandler)
	http.ListenAndServe(cfg.Http.Port, nil)

}
