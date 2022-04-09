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
	GetAll() ([] model.Ram)
	GetById(id uuid.UUID) (model.Ram, error)
	GetByName(name string) (model.Ram, error)
	Create(ram model.Ram) error
	Update(ramDTO dto.RamDTO) error
	Delete(id uuid.UUID) error
}

func NewRamServiceService(ramRepository repository.IRamRepository) IRamService {
	return &ramService{IRamRepository: ramRepository}
}

func (ramService *ramService) GetAll() []model.Ram {
	return ramService.IRamRepository.GetAll()
}

func (ramService *ramService) GetById(id uuid.UUID) (model.Ram, error) {
	return ramService.IRamRepository.GetById(id)
}

func (ramService *ramService) GetByName(name string) (model.Ram, error) {
	return ramService.IRamRepository.GetByName(name)
}

func (ramService *ramService) Create(ram model.Ram) error {
	ram.Id = uuid.New()
	return ramService.IRamRepository.Create(ram)
}

func (ramService *ramService) Update(ramDTO dto.RamDTO) error {
	ram, err := ramService.GetById(ramDTO.Id)
	if err != nil {
		return err
	}
	updatedRam := mapper.ToRam(ramDTO)
	updatedRam.Id = ram.Id
	return ramService.IRamRepository.Update(updatedRam)
}

func (ramService *ramService) Delete(id uuid.UUID) error {
	ram, err := ramService.GetById(id)
	if err != nil {
		return err
	}
	return ramService.IRamRepository.Delete(ram)
}