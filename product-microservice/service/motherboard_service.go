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

type motherboardService struct {
	IMotherboardRepository repository.IMotherboardRepository
}

type IMotherboardService interface {
	GetAll(page int, pageSize int) ([] model.Motherboard)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.Motherboard, error)
	SearchByName(page int, pageSize int, name string) ([]model.Motherboard, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.MotherboardFilter) ([]model.Motherboard, error)
	GetNumberOfRecordsFilter(filter filter.MotherboardFilter) int64
	GetManufacturers() []string
	GetProcessorTypes() []string
	GetSockets() []string
	GetFormFactors() []string
	Create(motherboard model.Motherboard) string
	Update(motherboardDTO dto.MotherboardDTO) string
	Delete(id uuid.UUID) error
}

func NewMotherboardService(motherboardRepository repository.IMotherboardRepository) IMotherboardService {
	return &motherboardService{IMotherboardRepository: motherboardRepository}
}

func (motherboardService *motherboardService) GetAll(page int, pageSize int) []model.Motherboard {
	return motherboardService.IMotherboardRepository.GetAll(page, pageSize)
}

func (motherboardService *motherboardService) GetNumberOfRecords() int64 {
	return motherboardService.IMotherboardRepository.GetNumberOfRecords()
}

func (motherboardService *motherboardService) GetById(id uuid.UUID) (model.Motherboard, error) {
	return motherboardService.IMotherboardRepository.GetById(id)
}

func (motherboardService *motherboardService) SearchByName(page int, pageSize int, name string) ([]model.Motherboard, error) {
	return motherboardService.IMotherboardRepository.SearchByName(page, pageSize, name)
}

func (motherboardService *motherboardService) GetNumberOfRecordsSearch(name string) int64 {
	return motherboardService.IMotherboardRepository.GetNumberOfRecordsSearch(name)
}

func (motherboardService *motherboardService) Filter(page int, pageSize int, filter filter.MotherboardFilter) ([]model.Motherboard, error) {
	return motherboardService.IMotherboardRepository.Filter(page, pageSize, filter)
}

func (motherboardService *motherboardService) GetNumberOfRecordsFilter(filter filter.MotherboardFilter) int64 {
	return motherboardService.IMotherboardRepository.GetNumberOfRecordsFilter(filter)
}

func (motherboardService *motherboardService) GetManufacturers() []string {
	return motherboardService.IMotherboardRepository.GetManufacturers()
}

func (motherboardService *motherboardService) GetProcessorTypes() []string {
	return motherboardService.IMotherboardRepository.GetProcessorTypes()
}

func (motherboardService *motherboardService) GetSockets() []string {
	return motherboardService.IMotherboardRepository.GetSockets()
}


func (motherboardService *motherboardService) GetFormFactors() []string {
	return motherboardService.IMotherboardRepository.GetFormFactors()
}

func (motherboardService *motherboardService) Create(motherboard model.Motherboard) string {
	msg := ""
	motherboard.Product.Id = uuid.New()
	motherboard.ProductId = motherboard.Product.Id
	motherboard.Product.Type = model.MOTHERBOARD
	err := motherboardService.IMotherboardRepository.Create(motherboard)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (motherboardService *motherboardService) Update(motherboardDTO dto.MotherboardDTO) string {
	msg := ""
	motherboard, err := motherboardService.GetById(motherboardDTO.Product.Id)
	if err != nil {
		return err.Error()
	}
	updatedMotherboard := mapper.ToMotherboard(motherboardDTO)
	updatedMotherboard.Product.Id = motherboard.Product.Id
	updatedMotherboard.ProductId = motherboard.Product.Id
	err = motherboardService.IMotherboardRepository.Update(updatedMotherboard)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (motherboardService *motherboardService) Delete(id uuid.UUID) error {
	motherboard, err := motherboardService.GetById(id)
	if err != nil {
		return err
	}
	return motherboardService.IMotherboardRepository.Delete(motherboard)
}