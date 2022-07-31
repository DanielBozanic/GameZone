//go:build wireinject
// +build wireinject

package di

import (
	"contact-and-report/api"
	"contact-and-report/repository"
	"contact-and-report/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitReportAPI(db *gorm.DB) api.ReportAPI {
	wire.Build(repository.NewReportRepository, service.NewReportService, api.NewReportAPI)
	return api.ReportAPI{}
}

func InitBanAPI(db *gorm.DB) api.BanAPI {
	wire.Build(repository.NewBanRepository, repository.NewReportRepository, service.NewBanService, api.NewBanAPI)
	return api.BanAPI{}
}

func InitContactAPI(db *gorm.DB) api.ContactAPI {
	wire.Build(repository.NewContactRepository, service.NewContactService, api.NewContactAPI)
	return api.ContactAPI{}
}