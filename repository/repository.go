package repository

import (
	"log"
	"time"

	"xiaoliuren/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	ds := sqlite.Open(config.DBPATH)
	db, err := gorm.Open(ds, &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		Db = db
	}

	sqlDB, _ := Db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
