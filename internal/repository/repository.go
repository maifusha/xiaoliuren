package repository

import (
	"time"

	"xiaoliuren/internal/config"
	"xiaoliuren/internal/util/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	ds := sqlite.Open(config.Conf.Sqlite3.Path)
	db, err := gorm.Open(ds, &gorm.Config{})
	if err != nil {
		logger.Fatalln(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalln(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	Db = db
}
