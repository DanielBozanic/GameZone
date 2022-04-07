//go:build wireinject
// +build wireinject

package di

import (
	"product/api"
	"product/repository"
	"product/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitVideoGameAPI(db *gorm.DB) api.VideoGameAPI {
	wire.Build(repository.NewVideoGameRepository, service.NewVideoGameService, api.NewVideoGameAPI)
	return api.VideoGameAPI{}
}

func InitConsoleAPI(db *gorm.DB) api.ConsoleAPI {
	wire.Build(repository.NewConsoleRepository, service.NewConsoleService, api.NewConsoleAPI)
	return api.ConsoleAPI{}
}