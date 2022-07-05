package service

import (
	"product/dto"
	"product/dto/filter"
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
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.PowerSupplyUnit, error)
	SearchByName(page int, pageSize int, name string) ([]model.PowerSupplyUnit, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.PowerSupplyUnitFilter) ([]model.PowerSupplyUnit, error)
	GetNumberOfRecordsFilter(filter filter.PowerSupplyUnitFilter) int64
	GetManufacturers() []string
	GetPowers() []string
	GetTypes() []string
	GetFormFactors() []string
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

func (powerSupplyUnitService *powerSupplyUnitService) GetNumberOfRecords() int64 {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetNumberOfRecords()
}

func (powerSupplyUnitService *powerSupplyUnitService) GetById(id uuid.UUID) (model.PowerSupplyUnit, error) {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetById(id)
}

func (powerSupplyUnitService *powerSupplyUnitService) SearchByName(page int, pageSize int, name string) ([]model.PowerSupplyUnit, error) {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.SearchByName(page, pageSize, name)
}

func (powerSupplyUnitService *powerSupplyUnitService) GetNumberOfRecordsSearch(name string) int64 {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetNumberOfRecordsSearch(name)
}

func (powerSupplyUnitService *powerSupplyUnitService) Filter(page int, pageSize int, filter filter.PowerSupplyUnitFilter) ([]model.PowerSupplyUnit, error) {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.Filter(page, pageSize, filter)
}

func (powerSupplyUnitService *powerSupplyUnitService) GetNumberOfRecordsFilter(filter filter.PowerSupplyUnitFilter) int64 {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetNumberOfRecordsFilter(filter)
}

func (powerSupplyUnitService *powerSupplyUnitService) GetManufacturers() []string {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetManufacturers()
}

func (powerSupplyUnitService *powerSupplyUnitService) GetPowers() []string {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetPowers()
}

func (powerSupplyUnitService *powerSupplyUnitService) GetTypes() []string {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetTypes()
}

func (powerSupplyUnitService *powerSupplyUnitService) GetFormFactors() []string {
	return powerSupplyUnitService.IPowerSupplyUnitRepository.GetFormFactors()
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