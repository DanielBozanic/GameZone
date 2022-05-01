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
	GetAll(page int, pageSize int) ([] model.PowerSupplyUnit)
	GetById(id uuid.UUID) (model.PowerSupplyUnit, error)
	SearchByName(page int, pageSize int, name string) ([]model.PowerSupplyUnit, error)
	Create(powerSupplyUnit model.PowerSupplyUnit) error
	Update(powerSupplyUnitDTO dto.PowerSupplyUnitDTO) error
	Delete(id uuid.UUID) error
}

func NewPowerSupplyUnitService(powerSupplyUnitRepository repository.IPowerSupplyUnitRepository) IPowerSupplyUnitService {
	return &powerSupplyUnitService{IPowerSupplyUnitRepository: powerSupplyUnitRepository}
}

func (powerSupplyUnitService *powerSupplyUnitService) GetAll(page int, pageSize int) []model.PowerSupplyUnit {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetAll(page, pageSize)
}

func (powerSupplyUnitService *powerSupplyUnitService) GetById(id uuid.UUID) (model.PowerSupplyUnit, error) {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetById(id)
}

func (powerSupplyUnitService *powerSupplyUnitService) SearchByName(page int, pageSize int, name string) ([]model.PowerSupplyUnit, error) {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.SearchByName(page, pageSize, name)
}

func (powerSupplyUnitService *powerSupplyUnitService) Create(powerSupplyUnit model.PowerSupplyUnit) error {
	powerSupplyUnit.Product.Id = uuid.New()
	powerSupplyUnit.ProductId = powerSupplyUnit.Product.Id
	powerSupplyUnit.Product.Type = model.POWER_SUPPLY_UNIT
	return powerSupplyUnitService.IPowerSupplyUnitRepository.Create(powerSupplyUnit)
}

func (powerSupplyUnitService *powerSupplyUnitService) Update(powerSupplyUnitDTO dto.PowerSupplyUnitDTO) error {
	powerSupplyUnit, err := powerSupplyUnitService.GetById(powerSupplyUnitDTO.Product.Id)
	if err != nil {
		return err
	}
	updatedPowerSupplyUnit := mapper.ToPowerSupplyUnit(powerSupplyUnitDTO)
	updatedPowerSupplyUnit.Product.Id = powerSupplyUnit.Product.Id
	updatedPowerSupplyUnit.ProductId = updatedPowerSupplyUnit.Product.Id
	return powerSupplyUnitService.IPowerSupplyUnitRepository.Update(updatedPowerSupplyUnit)
}

func (powerSupplyUnitService *powerSupplyUnitService) Delete(id uuid.UUID) error {
	powerSupplyUnit, err := powerSupplyUnitService.GetById(id)
	if err != nil {
		return err
	}
	return powerSupplyUnitService.IPowerSupplyUnitRepository.Delete(powerSupplyUnit)
}