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

func InitMotherboardAPI(db *gorm.DB) api.MotherboardAPI {
	wire.Build(repository.NewMotherboardRepository, service.NewMotherboardService, api.NewMotherboardAPI)
	return api.MotherboardAPI{}
}

func InitRamAPI(db *gorm.DB) api.RamAPI {
	wire.Build(repository.NewRamRepository, service.NewRamServiceService, api.NewRamAPI)
	return api.RamAPI{}
}

func InitSoldiStateDriveAPI(db *gorm.DB) api.SolidStateDriveAPI {
	wire.Build(repository.NewSolidStateDriveRepository, service.NewSolidStateDriveService, api.NewSolidStateDriveAPI)
	return api.SolidStateDriveAPI{}
}

func InitHardDiskDriveAPI(db *gorm.DB) api.HardDiskDriveAPI {
	wire.Build(repository.NewHardDiskDriveRepository, service.NewHardDiskDriveService, api.NewHardDiskDriveAPI)
	return api.HardDiskDriveAPI{}
}