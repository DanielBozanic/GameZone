package service

import (
	"errors"
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type processorService struct {
	IProcessorRepository repository.IProcessorRepository
}

type IProcessorService interface {
	GetAll(page int, pageSize int) ([] model.Processor)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.Processor, error)
	SearchByName(page int, pageSize int, name string) ([]model.Processor, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.ProcessorFilter) ([]model.Processor, error)
	GetNumberOfRecordsFilter(filter filter.ProcessorFilter) int64
	GetManufacturers() []string
	GetTypes() []string
	GetSockets() []string
	GetNumberOfCores() []uint
	GetThreads() []uint
	Create(processor model.Processor) string
	Update(processorDTO dto.ProcessorDTO) string
	Delete(id uuid.UUID) error
}

func NewProcessorService(processorRepository repository.IProcessorRepository) IProcessorService {
	return &processorService{IProcessorRepository: processorRepository}
}

func (processorService *processorService) GetAll(page int, pageSize int) []model.Processor {
	return processorService.IProcessorRepository.GetAll(page, pageSize)
}

func (processorService *processorService) GetNumberOfRecords() int64 {
	return processorService.IProcessorRepository.GetNumberOfRecords()
}

func (processorService *processorService) GetById(id uuid.UUID) (model.Processor, error) {
	return processorService.IProcessorRepository.GetById(id)
}

func (processorService *processorService) SearchByName(page int, pageSize int, name string) ([]model.Processor, error) {
	return processorService.IProcessorRepository.SearchByName(page, pageSize, name)
}

func (processorService *processorService) GetNumberOfRecordsSearch(name string) int64 {
	return processorService.IProcessorRepository.GetNumberOfRecordsSearch(name)
}

func (processorService *processorService) Filter(page int, pageSize int, filter filter.ProcessorFilter) ([]model.Processor, error) {
	return processorService.IProcessorRepository.Filter(page, pageSize, filter)
}

func (processorService *processorService) GetNumberOfRecordsFilter(filter filter.ProcessorFilter) int64 {
	return processorService.IProcessorRepository.GetNumberOfRecordsFilter(filter)
}

func (processorService *processorService) GetManufacturers() []string {
	return processorService.IProcessorRepository.GetManufacturers()
}

func (processorService *processorService) GetTypes() []string {
	return processorService.IProcessorRepository.GetManufacturers()
}

func (processorService *processorService) GetSockets() []string {
	return processorService.IProcessorRepository.GetSockets()
}


func (processorService *processorService) GetNumberOfCores() []uint {
	return processorService.IProcessorRepository.GetNumberOfCores()
}

func (processorService *processorService) GetThreads() []uint {
	return processorService.IProcessorRepository.GetThreads()
}

func (processorService *processorService) Create(processor model.Processor) string {
	msg := ""
	processor.Product.Id = uuid.New()
	processor.ProductId = processor.Product.Id
	processor.Product.Type = model.PROCESSOR
	err := processorService.IProcessorRepository.Create(processor)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (processorService *processorService) Update(processorDTO dto.ProcessorDTO) string {
	msg := ""
	processor, err := processorService.GetById(processorDTO.Product.Id)
	if err != nil {
		return err.Error()
	}
	updatedProcessor := mapper.ToProcessor(processorDTO)
	updatedProcessor.Product.Id = processor.Product.Id
	updatedProcessor.ProductId = processor.Product.Id
	err = processorService.IProcessorRepository.Update(updatedProcessor)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (processorService *processorService) Delete(id uuid.UUID) error {
	processor, err := processorService.GetById(id)
	if err != nil {
		return err
	}
	return processorService.IProcessorRepository.Delete(processor)
}