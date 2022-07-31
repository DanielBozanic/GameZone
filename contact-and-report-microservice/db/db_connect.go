package db

import (
	"contact-and-report/config"
	"contact-and-report/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	db, dbErr := gorm.Open(mysql.Open(cfg.DBURL), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(
		&model.Report{},
		&model.Ban{},
		&model.ContactMessage{},
	)

	return db, dbErr
}