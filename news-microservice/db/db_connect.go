package db

import (
	"news/config"
	"news/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	db, dbErr := gorm.Open(mysql.Open(cfg.DBURL), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(
		&model.NewsArticle{},
		&model.NewsComment{}, 
		&model.NewsSubscription{},
	)

	return db, dbErr
}