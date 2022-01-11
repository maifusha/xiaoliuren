package repository

import (
	"log"

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
}
