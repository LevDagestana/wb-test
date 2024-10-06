package db

import (
	"database/sql"
	"fmt"
	"log"
	"wb/config"
	"wb/logger"
)

var Db *sql.DB

func InitDb(cfg config.DatabaseConfig) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	var err error

	Db, err = sql.Open(cfg.DriverName, connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		logger.Log.WithError(err).Fatal("Ошибка подключения к БД:")

	}

	logger.Log.Info("Подключение к БД установлено")
}
