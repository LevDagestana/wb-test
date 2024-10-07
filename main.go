package main

import (
	"net/http"

	"wb/handlers"
	"wb/kafka"
	"wb/repository"

	"wb/db"

	"wb/config"

	"wb/logger"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	logger.InitLogger()
	router := gin.Default()

	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Log.WithError(err).Fatal("Ошибка загрузки конфигурации:")
	}
	db.InitDb(cfg)
	kafka.InitKafka(cfg)
	repository.LoadCache()
	router.StaticFS("/static", http.Dir("./static"))

	router.GET("/order/:id", handlers.GetOrderByIdHandler)
	router.Run(cfg.Port)

}
