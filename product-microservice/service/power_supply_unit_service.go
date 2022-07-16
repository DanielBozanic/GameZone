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
	Create(powerSupplyUnit model.PowerSupplyUnit) string
	Update(powerSupplyUnitDTO dto.PowerSupplyUnitDTO) string
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

func (powerSupplyUnitService *powerSupplyUnitService) Create(powerSupplyUnit model.PowerSupplyUnit) string {
	msg := ""
	powerSupplyUnit.Product.Id = uuid.New()
	powerSupplyUnit.ProductId = powerSupplyUnit.Product.Id
	powerSupplyUnit.Product.Type = model.POWER_SUPPLY_UNIT
	err := powerSupplyUnitService.IPowerSupplyUnitRepository.Create(powerSupplyUnit)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (powerSupplyUnitService *powerSupplyUnitService) Update(powerSupplyUnitDTO dto.PowerSupplyUnitDTO) string {
	msg := ""
	powerSupplyUnit, err := powerSupplyUnitService.GetById(powerSupplyUnitDTO.Product.Id)
	if err != nil {
		return err.Error()
	}
	updatedPowerSupplyUnit := mapper.ToPowerSupplyUnit(powerSupplyUnitDTO)
	updatedPowerSupplyUnit.Product.Id = powerSupplyUnit.Product.Id
	updatedPowerSupplyUnit.ProductId = updatedPowerSupplyUnit.Product.Id
	err = powerSupplyUnitService.IPowerSupplyUnitRepository.Update(updatedPowerSupplyUnit)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (powerSupplyUnitService *powerSupplyUnitService) Delete(id uuid.UUID) error {
	powerSupplyUnit, err := powerSupplyUnitService.GetById(id)
	if err != nil {
		return err
	}
	return powerSupplyUnitService.IPowerSupplyUnitRepository.Delete(powerSupplyUnit)
}