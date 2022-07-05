package service

import (
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type solidStateDriveService struct {
	ISolidStateDriveRepository repository.ISolidStateDriveRepository
}

type ISolidStateDriveService interface {
	GetAll(page int, pageSize int) ([] model.SolidStateDrive)
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.SolidStateDrive, error)
	SearchByName(page int, pageSize int, name string) ([]model.SolidStateDrive, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.SolidStateDriveFilter) ([]model.SolidStateDrive, error)
	GetNumberOfRecordsFilter(filter filter.SolidStateDriveFilter) int64
	GetCapacities() []string
	GetForms() []string
	GetManufacturers() []string
	GetMaxSequentialReads() []string
	GetMaxSequentialWrites() []string
	Create(solidStateDrive model.SolidStateDrive) error
	Update(solidStateDriveDTO dto.SolidStateDriveDTO) error
	Delete(id uuid.UUID) error
}

func NewSolidStateDriveService(solidStateDriveRepository repository.ISolidStateDriveRepository) ISolidStateDriveService {
	return &solidStateDriveService{ISolidStateDriveRepository: solidStateDriveRepository}
}

func (solidStateDriveService *solidStateDriveService) GetAll(page int, pageSize int) []model.SolidStateDrive {
	return solidStateDriveService.ISolidStateDriveRepository.GetAll(page, pageSize)
}

func (solidStateDriveService *solidStateDriveService) GetNumberOfRecords() int64 {
	return solidStateDriveService.ISolidStateDriveRepository.GetNumberOfRecords()
}

func (solidStateDriveService *solidStateDriveService) GetById(id uuid.UUID) (model.SolidStateDrive, error) {
	return solidStateDriveService.ISolidStateDriveRepository.GetById(id)
}

func (solidStateDriveService *solidStateDriveService) SearchByName(page int, pageSize int, name string) ([]model.SolidStateDrive, error) {
	return solidStateDriveService.ISolidStateDriveRepository.SearchByName(page, pageSize, name)
}

func (solidStateDriveService *solidStateDriveService) GetNumberOfRecordsSearch(name string) int64 {
	return solidStateDriveService.ISolidStateDriveRepository.GetNumberOfRecordsSearch(name)
}

func (solidStateDriveService *solidStateDriveService) Filter(page int, pageSize int, filter filter.SolidStateDriveFilter) ([]model.SolidStateDrive, error) {
	return solidStateDriveService.ISolidStateDriveRepository.Filter(page, pageSize, filter)
}

func (solidStateDriveService *solidStateDriveService) GetNumberOfRecordsFilter(filter filter.SolidStateDriveFilter) int64 {
	return solidStateDriveService.ISolidStateDriveRepository.GetNumberOfRecordsFilter(filter)
}

func (solidStateDriveService *solidStateDriveService) GetCapacities() []string {
	return solidStateDriveService.ISolidStateDriveRepository.GetCapacities()
}

func (solidStateDriveService *solidStateDriveService) GetForms() []string {
	return solidStateDriveService.ISolidStateDriveRepository.GetForms()
}

func (solidStateDriveService *solidStateDriveService) GetManufacturers() []string {
	return solidStateDriveService.ISolidStateDriveRepository.GetManufacturers()
}

func (solidStateDriveService *solidStateDriveService) GetMaxSequentialReads() []string {
	return solidStateDriveService.ISolidStateDriveRepository.GetMaxSequentialReads()
}

func (solidStateDriveService *solidStateDriveService) GetMaxSequentialWrites() []string {
	return solidStateDriveService.ISolidStateDriveRepository.GetMaxSequentialWrites()
}

func (solidStateDriveService *solidStateDriveService) Create(solidStateDrive model.SolidStateDrive) error {
	solidStateDrive.Product.Id = uuid.New()
	solidStateDrive.ProductId = solidStateDrive.Product.Id
	solidStateDrive.Product.Type = model.SOLID_STATE_DRIVE
	return solidStateDriveService.ISolidStateDriveRepository.Create(solidStateDrive)
}

func (solidStateDriveService *solidStateDriveService) Update(solidStateDriveDTO dto.SolidStateDriveDTO) error {
	solidStateDrive, err := solidStateDriveService.GetById(solidStateDriveDTO.Product.Id)
	if err != nil {
		return err
	}
	updatedSolidStateDrive := mapper.ToSolidStateDrive(solidStateDriveDTO)
	updatedSolidStateDrive.Product.Id = solidStateDrive.Product.Id
	updatedSolidStateDrive.ProductId = solidStateDrive.Product.Id
	return solidStateDriveService.ISolidStateDriveRepository.Update(updatedSolidStateDrive)
}

func (solidStateDriveService *solidStateDriveService) Delete(id uuid.UUID) error {
	solidStateDrive, err := solidStateDriveService.GetById(id)
	if err != nil {
		return err
	}
	return solidStateDriveService.ISolidStateDriveRepository.Delete(solidStateDrive)
}