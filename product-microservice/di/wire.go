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

func InitGraphicsCardAPI(db *gorm.DB) api.GraphicsCardAPI {
	wire.Build(repository.NewGraphicsCardRepository, service.NewGraphicsCardService, api.NewGraphicsCardAPI)
	return api.GraphicsCardAPI{}
}

func InitProcessorAPI(db *gorm.DB) api.ProcessorAPI {
	wire.Build(repository.NewProcessorRepository, service.NewProcessorServiceService, api.NewProcessorAPI)
	return api.ProcessorAPI{}
}