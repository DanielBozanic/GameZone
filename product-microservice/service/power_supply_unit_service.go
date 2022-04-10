package service

import (
	"product/dto"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type powerSupplyUnitService struct {
	IPowerSupplyUnitRepository repository.IPowerSupplyUnitRepository
}

type IPowerSupplyUnitService interface {
	GetAll() ([] model.PowerSupplyUnit)
	GetById(id uuid.UUID) (model.PowerSupplyUnit, error)
	GetByName(name string) (model.PowerSupplyUnit, error)
	Create(powerSupplyUnit model.PowerSupplyUnit) error
	Update(powerSupplyUnitDTO dto.PowerSupplyUnitDTO) error
	Delete(id uuid.UUID) error
}

func NewPowerSupplyUnitService(powerSupplyUnitRepository repository.IPowerSupplyUnitRepository) IPowerSupplyUnitService {
	return &powerSupplyUnitService{IPowerSupplyUnitRepository: powerSupplyUnitRepository}
}

func (powerSupplyUnitService *powerSupplyUnitService) GetAll() []model.PowerSupplyUnit {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetAll()
}

func (powerSupplyUnitService *powerSupplyUnitService) GetById(id uuid.UUID) (model.PowerSupplyUnit, error) {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetById(id)
}

func (powerSupplyUnitService *powerSupplyUnitService) GetByName(name string) (model.PowerSupplyUnit, error) {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetByName(name)
}

func (powerSupplyUnitService *powerSupplyUnitService) Create(powerSupplyUnit model.PowerSupplyUnit) error {
	powerSupplyUnit.Id = uuid.New()
	return powerSupplyUnitService.IPowerSupplyUnitRepository.Create(powerSupplyUnit)
}

func (powerSupplyUnitService *powerSupplyUnitService) Update(powerSupplyUnitDTO dto.PowerSupplyUnitDTO) error {
	powerSupplyUnit, err := powerSupplyUnitService.GetById(powerSupplyUnitDTO.Id)
	if err != nil {
		return err
	}
	updatedPowerSupplyUnit := mapper.ToPowerSupplyUnit(powerSupplyUnitDTO)
	updatedPowerSupplyUnit.Id = powerSupplyUnit.Id
	return powerSupplyUnitService.IPowerSupplyUnitRepository.Update(updatedPowerSupplyUnit)
}

func (powerSupplyUnitService *powerSupplyUnitService) Delete(id uuid.UUID) error {
	powerSupplyUnit, err := powerSupplyUnitService.GetById(id)
	if err != nil {
		return err
	}
	return powerSupplyUnitService.IPowerSupplyUnitRepository.Delete(powerSupplyUnit)
}