package db

import (
	"product/config"
	"product/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	db, dbErr := gorm.Open(mysql.Open(cfg.DBURL), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&model.VideoGame{})

	return db, dbErr
}