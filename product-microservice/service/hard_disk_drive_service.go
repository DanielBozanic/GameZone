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
	GetAll(page int, pageSize int) ([] model.HardDiskDrive)
	GetById(id uuid.UUID) (model.HardDiskDrive, error)
	SearchByName(page int, pageSize int, name string) ([]model.HardDiskDrive, error)
	Create(hardDiskDrive model.HardDiskDrive) error
	Update(hardDiskDriveDTO dto.HardDiskDriveDTO) error
	Delete(id uuid.UUID) error
}

func NewHardDiskDriveService(hardDiskDriveRepository repository.IHardDiskDriveRepository) IHardDiskDriveService {
	return &hardDiskDriveService{IHardDiskDriveRepository: hardDiskDriveRepository}
}

func (hardDiskDriveService *hardDiskDriveService) GetAll(page int, pageSize int) []model.HardDiskDrive {
	return hardDiskDriveService.IHardDiskDriveRepository.GetAll(page, pageSize)
}

func (hardDiskDriveService *hardDiskDriveService) GetById(id uuid.UUID) (model.HardDiskDrive, error) {
	return hardDiskDriveService.IHardDiskDriveRepository.GetById(id)
}

func (hardDiskDriveService *hardDiskDriveService) SearchByName(page int, pageSize int, name string) ([]model.HardDiskDrive, error) {
	return hardDiskDriveService.IHardDiskDriveRepository.SearchByName(page, pageSize, name)
}

func (hardDiskDriveService *hardDiskDriveService) Create(hardDiskDrive model.HardDiskDrive) error {
	hardDiskDrive.Product.Id = uuid.New()
	hardDiskDrive.ProductId = hardDiskDrive.Product.Id
	hardDiskDrive.Product.Type = model.HARD_DISK_DRIVE
	return hardDiskDriveService.IHardDiskDriveRepository.Create(hardDiskDrive)
}

func (hardDiskDriveService *hardDiskDriveService) Update(hardDiskDriveDTO dto.HardDiskDriveDTO) error {
	hardDiskDrive, err := hardDiskDriveService.GetById(hardDiskDriveDTO.Product.Id)
	if err != nil {
		return err
	}
	updatedHardDiskDrive := mapper.ToHardDiskDrive(hardDiskDriveDTO)
	updatedHardDiskDrive.Product.Id = hardDiskDrive.Product.Id
	updatedHardDiskDrive.ProductId = hardDiskDrive.Product.Id
	return hardDiskDriveService.IHardDiskDriveRepository.Update(updatedHardDiskDrive)
}

func (hardDiskDriveService *hardDiskDriveService) Delete(id uuid.UUID) error {
	hardDiskDrive, err := hardDiskDriveService.GetById(id)
	if err != nil {
		return err
	}
	return hardDiskDriveService.IHardDiskDriveRepository.Delete(hardDiskDrive)
}