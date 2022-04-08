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
	GetAll() ([] model.Processor)
	GetById(id uuid.UUID) (model.Processor, error)
	GetByName(name string) (model.Processor, error)
	Create(processor model.Processor) error
	Update(processor model.Processor) error
	Delete(processor model.Processor) error
}

func NewProcessorRepository(DB *gorm.DB) IProcessorRepository {
	return &processorRepository{Database: DB}
}

func (processorRepo *processorRepository) GetAll() []model.Processor {
	var processors []model.Processor
	processorRepo.Database.Find(&processors)
	return processors
}

func (processorRepo *processorRepository) GetById(id uuid.UUID) (model.Processor, error) {
	var processor model.Processor
	result := processorRepo.Database.First(&processor, id)
	return processor, result.Error
}

func (processorRepo *processorRepository) GetByName(name string) (model.Processor, error) {
	var processor model.Processor
	result := processorRepo.Database.Find(&processor, "name = ?", name)
	return processor, result.Error
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