package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type monitorRepository struct {
	Database *gorm.DB
}

type IMonitorRepository interface {
	GetAll() ([] model.Monitor)
	GetById(id uuid.UUID) (model.Monitor, error)
	GetByName(name string) (model.Monitor, error)
	Create(monitor model.Monitor) error
	Update(monitor model.Monitor) error
	Delete(monitor model.Monitor) error
}

func NewMonitorRepository(DB *gorm.DB) IMonitorRepository {
	return &monitorRepository{Database: DB}
}

func (monitorRepo *monitorRepository) GetAll() []model.Monitor {
	var monitors []model.Monitor
	monitorRepo.Database.Preload("Product").Find(&monitors)
	return monitors
}

func (monitorRepo *monitorRepository) GetById(id uuid.UUID) (model.Monitor, error) {
	var monitor model.Monitor
	result := monitorRepo.Database.Preload("Product").First(&monitor, id)
	return monitor, result.Error
}

func (monitorRepo *monitorRepository) GetByName(name string) (model.Monitor, error) {
	var monitor model.Monitor
	result := monitorRepo.Database.Preload("Product").Find(&monitor, "name = ?", name)
	return monitor, result.Error
}

func (monitorRepo *monitorRepository) Create(monitor model.Monitor) error {
	result := monitorRepo.Database.Create(&monitor)
	return result.Error
}

func (monitorRepo *monitorRepository) Update(monitor model.Monitor) error {
	result := monitorRepo.Database.Save(&monitor)
	return result.Error
}

func (monitorRepo *monitorRepository) Delete(monitor model.Monitor) error {
	result := monitorRepo.Database.Delete(&monitor)
	return result.Error
}