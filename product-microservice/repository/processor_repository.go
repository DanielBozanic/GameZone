package repository

import (
	"product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type processorRepository struct {
	Database *gorm.DB
}

type IProcessorRepository interface {
	GetAll(page int, pageSize int) ([] model.Processor)
	GetById(id uuid.UUID) (model.Processor, error)
	SearchByName(page int, pageSize int, name string) ([]model.Processor, error)
	Create(processor model.Processor) error
	Update(processor model.Processor) error
	Delete(processor model.Processor) error
}

func NewProcessorRepository(DB *gorm.DB) IProcessorRepository {
	return &processorRepository{Database: DB}
}

func (processorRepo *processorRepository) GetAll(page int, pageSize int) []model.Processor {
	var processors []model.Processor
	offset := (page - 1) * pageSize
	processorRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Find(&processors)
	return processors
}

func (processorRepo *processorRepository) GetById(id uuid.UUID) (model.Processor, error) {
	var processor model.Processor
	result := processorRepo.Database.Preload("Product").First(&processor, id)
	return processor, result.Error
}

func (processorRepo *processorRepository) SearchByName(page int, pageSize int, name string) ([]model.Processor, error) {
	var processors []model.Processor
	offset := (page - 1) * pageSize
	result := processorRepo.Database.
		Offset(offset).Limit(pageSize).
		Preload("Product").
		Joins("JOIN products ON products.id = processors.product_id").
		Where("products.name LIKE ?", "%" + name + "%").
		Find(&processors)
	return processors, result.Error
}

func (processorRepo *processorRepository) Create(processor model.Processor) error {
	result := processorRepo.Database.Create(&processor)
	return result.Error
}

func (processorRepo *processorRepository) Update(processor model.Processor) error {
	result := processorRepo.Database.Save(&processor)
	return result.Error
}

func (processorRepo *processorRepository) Delete(processor model.Processor) error {
	result := processorRepo.Database.Delete(&processor)
	return result.Error
}