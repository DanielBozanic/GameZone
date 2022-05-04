package service

import (
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type motherboardService struct {
	IMotherboardRepository repository.IMotherboardRepository
}

type IMotherboardService interface {
	GetAll(page int, pageSize int) ([] model.Motherboard)
	GetById(id uuid.UUID) (model.Motherboard, error)
	SearchByName(page int, pageSize int, name string) ([]model.Motherboard, error)
	Filter(page int, pageSize int, filter filter.MotherboardFilter) ([]model.Motherboard, error)
	GetManufacturers() []string
	GetProcessorTypes() []string
	GetSockets() []string
	GetFormFactors() []string
	Create(motherboard model.Motherboard) error
	Update(motherboardDTO dto.MotherboardDTO) error
	Delete(id uuid.UUID) error
}

func NewMotherboardService(motherboardRepository repository.IMotherboardRepository) IMotherboardService {
	return &motherboardService{IMotherboardRepository: motherboardRepository}
}

func (motherboardService *motherboardService) GetAll(page int, pageSize int) []model.Motherboard {
	return motherboardService.IMotherboardRepository.GetAll(page, pageSize)
}

func (motherboardService *motherboardService) GetById(id uuid.UUID) (model.Motherboard, error) {
	return motherboardService.IMotherboardRepository.GetById(id)
}

func (motherboardService *motherboardService) SearchByName(page int, pageSize int, name string) ([]model.Motherboard, error) {
	return motherboardService.IMotherboardRepository.SearchByName(page, pageSize, name)
}

func (motherboardService *motherboardService) Filter(page int, pageSize int, filter filter.MotherboardFilter) ([]model.Motherboard, error) {
	return motherboardService.IMotherboardRepository.Filter(page, pageSize, filter)
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

func (motherboardService *motherboardService) Create(motherboard model.Motherboard) error {
	motherboard.Product.Id = uuid.New()
	motherboard.ProductId = motherboard.Product.Id
	motherboard.Product.Type = model.MOTHERBOARD
	return motherboardService.IMotherboardRepository.Create(motherboard)
}

func (motherboardService *motherboardService) Update(motherboardDTO dto.MotherboardDTO) error {
	motherboard, err := motherboardService.GetById(motherboardDTO.Product.Id)
	if err != nil {
		return err
	}
	updatedMotherboard := mapper.ToMotherboard(motherboardDTO)
	updatedMotherboard.Product.Id = motherboard.Product.Id
	updatedMotherboard.ProductId = motherboard.Product.Id
	return motherboardService.IMotherboardRepository.Update(updatedMotherboard)
}

func (motherboardService *motherboardService) Delete(id uuid.UUID) error {
	motherboard, err := motherboardService.GetById(id)
	if err != nil {
		return err
	}
	return motherboardService.IMotherboardRepository.Delete(motherboard)
}