// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"contact-and-report/api"
	"contact-and-report/repository"
	"contact-and-report/service"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitReportAPI(db *gorm.DB) api.ReportAPI {
	iReportRepository := repository.NewReportRepository(db)
	iReportService := service.NewReportService(iReportRepository)
	reportAPI := api.NewReportAPI(iReportService)
	return reportAPI
}

func InitBanAPI(db *gorm.DB) api.BanAPI {
	iBanRepository := repository.NewBanRepository(db)
	iBanService := service.NewBanService(iBanRepository)
	banAPI := api.NewBanAPI(iBanService)
	return banAPI
}

func InitContactAPI(db *gorm.DB) api.ContactAPI {
	iContactRepository := repository.NewContactRepository(db)
	iContactService := service.NewContactService(iContactRepository)
	contactAPI := api.NewContactAPI(iContactService)
	return contactAPI
}
