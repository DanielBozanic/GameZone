package service

import (
	"product/dto"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type hardDiskDriveService struct {
	IHardDiskDriveRepository repository.IHardDiskDriveRepository
}

type IHardDiskDriveService interface {
	GetAll() ([] model.HardDiskDrive)
	GetById(id uuid.UUID) (model.HardDiskDrive, error)
	GetByName(name string) (model.HardDiskDrive, error)
	Create(hardDiskDrive model.HardDiskDrive) error
	Update(hardDiskDriveDTO dto.HardDiskDriveDTO) error
	Delete(id uuid.UUID) error
}

func NewHardDiskDriveService(hardDiskDriveRepository repository.IHardDiskDriveRepository) IHardDiskDriveService {
	return &hardDiskDriveService{IHardDiskDriveRepository: hardDiskDriveRepository}
}

func (hardDiskDriveService *hardDiskDriveService) GetAll() []model.HardDiskDrive {
	return hardDiskDriveService.IHardDiskDriveRepository.GetAll()
}

func (hardDiskDriveService *hardDiskDriveService) GetById(id uuid.UUID) (model.HardDiskDrive, error) {
	return hardDiskDriveService.IHardDiskDriveRepository.GetById(id)
}

func (hardDiskDriveService *hardDiskDriveService) GetByName(name string) (model.HardDiskDrive, error) {
	return hardDiskDriveService.IHardDiskDriveRepository.GetByName(name)
}

func (hardDiskDriveService *hardDiskDriveService) Create(hardDiskDrive model.HardDiskDrive) error {
	hardDiskDrive.Id = uuid.New()
	return hardDiskDriveService.IHardDiskDriveRepository.Create(hardDiskDrive)
}

func (hardDiskDriveService *hardDiskDriveService) Update(hardDiskDriveDTO dto.HardDiskDriveDTO) error {
	hardDiskDrive, err := hardDiskDriveService.GetById(hardDiskDriveDTO.Id)
	if err != nil {
		return err
	}
	updatedHardDiskDrive := mapper.ToHardDiskDrive(hardDiskDriveDTO)
	updatedHardDiskDrive.Id = hardDiskDrive.Id
	return hardDiskDriveService.IHardDiskDriveRepository.Update(updatedHardDiskDrive)
}

func (hardDiskDriveService *hardDiskDriveService) Delete(id uuid.UUID) error {
	hardDiskDrive, err := hardDiskDriveService.GetById(id)
	if err != nil {
		return err
	}
	return hardDiskDriveService.IHardDiskDriveRepository.Delete(hardDiskDrive)
}