package service

import (
	"product/dto"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type processorService struct {
	IProcessorRepository repository.IProcessorRepository
}

type IProcessorService interface {
	GetAll() ([] model.Processor)
	GetById(id uuid.UUID) (model.Processor, error)
	GetByName(name string) (model.Processor, error)
	Create(processor model.Processor) error
	Update(processorDTO dto.ProcessorDTO) error
	Delete(id uuid.UUID) error
}

func NewProcessorService(processorRepository repository.IProcessorRepository) IProcessorService {
	return &processorService{IProcessorRepository: processorRepository}
}

func (processorService *processorService) GetAll() []model.Processor {
	return processorService.IProcessorRepository.GetAll()
}

func (processorService *processorService) GetById(id uuid.UUID) (model.Processor, error) {
	return processorService.IProcessorRepository.GetById(id)
}

func (processorService *processorService) GetByName(name string) (model.Processor, error) {
	return processorService.IProcessorRepository.GetByName(name)
}

func (processorService *processorService) Create(processor model.Processor) error {
	processor.Product.Id = uuid.New()
	processor.ProductId = processor.Product.Id
	processor.Product.Type = model.PROCESSOR
	return processorService.IProcessorRepository.Create(processor)
}

func (processorService *processorService) Update(processorDTO dto.ProcessorDTO) error {
	processor, err := processorService.GetById(processorDTO.Product.Id)
	if err != nil {
		return err
	}
	updatedProcessor := mapper.ToProcessor(processorDTO)
	updatedProcessor.Product.Id = processor.Product.Id
	updatedProcessor.ProductId = processor.Product.Id
	return processorService.IProcessorRepository.Update(updatedProcessor)
}

func (processorService *processorService) Delete(id uuid.UUID) error {
	processor, err := processorService.GetById(id)
	if err != nil {
		return err
	}
	return processorService.IProcessorRepository.Delete(processor)
}