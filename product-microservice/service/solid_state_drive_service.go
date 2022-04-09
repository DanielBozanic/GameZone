package service

import (
	"product/dto"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type solidStateDriveService struct {
	ISolidStateDriveRepository repository.ISolidStateDriveRepository
}

type ISolidStateDriveService interface {
	GetAll() ([] model.SolidStateDrive)
	GetById(id uuid.UUID) (model.SolidStateDrive, error)
	GetByName(name string) (model.SolidStateDrive, error)
	Create(solidStateDrive model.SolidStateDrive) error
	Update(solidStateDriveDTO dto.SolidStateDriveDTO) error
	Delete(id uuid.UUID) error
}

func NewSolidStateDriveService(solidStateDriveRepository repository.ISolidStateDriveRepository) ISolidStateDriveService {
	return &solidStateDriveService{ISolidStateDriveRepository: solidStateDriveRepository}
}

func (solidStateDriveService *solidStateDriveService) GetAll() []model.SolidStateDrive {
	return solidStateDriveService.ISolidStateDriveRepository.GetAll()
}

func (solidStateDriveService *solidStateDriveService) GetById(id uuid.UUID) (model.SolidStateDrive, error) {
	return solidStateDriveService.ISolidStateDriveRepository.GetById(id)
}

func (solidStateDriveService *solidStateDriveService) GetByName(name string) (model.SolidStateDrive, error) {
	return solidStateDriveService.ISolidStateDriveRepository.GetByName(name)
}

func (solidStateDriveService *solidStateDriveService) Create(solidStateDrive model.SolidStateDrive) error {
	solidStateDrive.Id = uuid.New()
	return solidStateDriveService.ISolidStateDriveRepository.Create(solidStateDrive)
}

func (solidStateDriveService *solidStateDriveService) Update(solidStateDriveDTO dto.SolidStateDriveDTO) error {
	solidStateDrive, err := solidStateDriveService.GetById(solidStateDriveDTO.Id)
	if err != nil {
		return err
	}
	updatedSolidStateDrive := mapper.ToSolidStateDrive(solidStateDriveDTO)
	updatedSolidStateDrive.Id = solidStateDrive.Id
	return solidStateDriveService.ISolidStateDriveRepository.Update(updatedSolidStateDrive)
}

func (solidStateDriveService *solidStateDriveService) Delete(id uuid.UUID) error {
	solidStateDrive, err := solidStateDriveService.GetById(id)
	if err != nil {
		return err
	}
	return solidStateDriveService.ISolidStateDriveRepository.Delete(solidStateDrive)
}