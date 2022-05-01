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
	GetAll(page int, pageSize int) ([] model.Monitor)
	GetById(id uuid.UUID) (model.Monitor, error)
	SearchByName(page int, pageSize int, name string) ([]model.Monitor, error)
	Create(monitor model.Monitor) error
	Update(monitor model.Monitor) error
	Delete(monitor model.Monitor) error
}

func NewMonitorRepository(DB *gorm.DB) IMonitorRepository {
	return &monitorRepository{Database: DB}
}

func (monitorRepo *monitorRepository) GetAll(page int, pageSize int) []model.Monitor {
	var monitors []model.Monitor
	offset := (page - 1) * pageSize
	monitorRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&monitors)
	return monitors
}

func (monitorRepo *monitorRepository) GetById(id uuid.UUID) (model.Monitor, error) {
	var monitor model.Monitor
	result := monitorRepo.Database.Preload("Product").First(&monitor, id)
	return monitor, result.Error
}

func (monitorRepo *monitorRepository) SearchByName(page int, pageSize int, name string) ([]model.Monitor, error) {
	var monitors []model.Monitor
	offset := (page - 1) * pageSize
	result := monitorRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = monitors.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&monitors)
	return monitors, result.Error
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