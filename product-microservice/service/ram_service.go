package service

import (
	"product/dto"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type ramService struct {
	IRamRepository repository.IRamRepository
}

type IRamService interface {
	GetAll(page int, pageSize int) ([] model.Ram)
	GetById(id uuid.UUID) (model.Ram, error)
	SearchByName(page int, pageSize int, name string) ([]model.Ram, error)
	Create(ram model.Ram) error
	Update(ramDTO dto.RamDTO) error
	Delete(id uuid.UUID) error
}

func NewRamServiceService(ramRepository repository.IRamRepository) IRamService {
	return &ramService{IRamRepository: ramRepository}
}

func (ramService *ramService) GetAll(page int, pageSize int) []model.Ram {
	return ramService.IRamRepository.GetAll(page, pageSize)
}

func (ramService *ramService) GetById(id uuid.UUID) (model.Ram, error) {
	return ramService.IRamRepository.GetById(id)
}

func (ramService *ramService) SearchByName(page int, pageSize int, name string) ([]model.Ram, error) {
	return ramService.IRamRepository.SearchByName(page, pageSize, name)
}

func (ramService *ramService) Create(ram model.Ram) error {
	ram.Product.Id = uuid.New()
	ram.ProductId = ram.Product.Id
	ram.Product.Type = model.RAM
	return ramService.IRamRepository.Create(ram)
}

func (ramService *ramService) Update(ramDTO dto.RamDTO) error {
	ram, err := ramService.GetById(ramDTO.Product.Id)
	if err != nil {
		return err
	}
	updatedRam := mapper.ToRam(ramDTO)
	updatedRam.Product.Id = ram.Product.Id
	updatedRam.ProductId = ram.Product.Id
	return ramService.IRamRepository.Update(updatedRam)
}

func (ramService *ramService) Delete(id uuid.UUID) error {
	ram, err := ramService.GetById(id)
	if err != nil {
		return err
	}
	return ramService.IRamRepository.Delete(ram)
}