package service

import (
	"product/dto"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type motherboardService struct {
	IMotherboardRepository repository.IMotherboardRepository
}

type IMotherboardService interface {
	GetAll() ([] model.Motherboard)
	GetById(id uuid.UUID) (model.Motherboard, error)
	GetByName(name string) (model.Motherboard, error)
	Create(motherboard model.Motherboard) error
	Update(motherboardDTO dto.MotherboardDTO) error
	Delete(id uuid.UUID) error
}

func NewMotherboardService(motherboardRepository repository.IMotherboardRepository) IMotherboardService {
	return &motherboardService{IMotherboardRepository: motherboardRepository}
}

func (motherboardService *motherboardService) GetAll() []model.Motherboard {
	return motherboardService.IMotherboardRepository.GetAll()
}

func (motherboardService *motherboardService) GetById(id uuid.UUID) (model.Motherboard, error) {
	return motherboardService.IMotherboardRepository.GetById(id)
}

func (motherboardService *motherboardService) GetByName(name string) (model.Motherboard, error) {
	return motherboardService.IMotherboardRepository.GetByName(name)
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