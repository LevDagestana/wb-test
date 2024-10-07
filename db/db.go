package db

import (
	"database/sql"
	"log"
	"wb/config"
	"wb/logger"
)

var Db *sql.DB

func InitDb(cfg *config.Config) {

	var err error

	Db, err = sql.Open(cfg.DbDriverName, cfg.DbScheme)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		logger.Log.WithError(err).Fatal("Ошибка подключения к БД:")

	}

	logger.Log.Info("Подключение к БД установлено")
}
