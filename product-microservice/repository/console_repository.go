package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type consoleRepository struct {
	Database *gorm.DB
}

type IConsoleRepository interface {
	GetAll() ([] model.Console)
	GetById(id uuid.UUID) (model.Console, error)
	GetByName(name string) (model.Console, error)
	Create(videoGame model.Console) error
	Update(videoGame model.Console) error
	Delete(videoGame model.Console) error
}

func NewConsoleRepository(DB *gorm.DB) IConsoleRepository {
	return &consoleRepository{Database: DB}
}

func (consoleRepo *consoleRepository) GetAll() []model.Console {
	var consoles []model.Console
	consoleRepo.Database.Preload("Product").Find(&consoles)
	return consoles
}

func (consoleRepo *consoleRepository) GetById(id uuid.UUID) (model.Console, error) {
	var console model.Console
	result := consoleRepo.Database.Preload("Product").First(&console, id)
	return console, result.Error
}

func (consoleRepo *consoleRepository) GetByName(name string) (model.Console, error) {
	var console model.Console
	result := consoleRepo.Database.Preload("Product").Find(&console, "name = ?", name)
	return console, result.Error
}

func (consoleRepo *consoleRepository) Create(console model.Console) error {
	result := consoleRepo.Database.Create(&console)
	return result.Error
}

func (consoleRepo *consoleRepository) Update(console model.Console) error {
	result := consoleRepo.Database.Save(&console)
	return result.Error
}

func (consoleRepo *consoleRepository) Delete(console model.Console) error {
	result := consoleRepo.Database.Delete(&console)
	return result.Error
}