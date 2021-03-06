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

	db.AutoMigrate(
		&model.Product{},
		&model.VideoGame{}, 
		&model.Console{}, 
		&model.GraphicsCard{}, 
		&model.Processor{}, 
		&model.Motherboard{},
		&model.Ram{},
		&model.SolidStateDrive{},
		&model.HardDiskDrive{},
		&model.Monitor{},
		&model.PowerSupplyUnit{},
		&model.Keyboard{},
		&model.Mouse{},
		&model.Headphones{},
		&model.ProductPurchase{},
		&model.ProductPurchaseDetail{},
		&model.File{},
		&model.ProductAlert{},
	)

	return db, dbErr
}