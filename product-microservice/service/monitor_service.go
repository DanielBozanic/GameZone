package service

import (
	"product/dto"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/google/uuid"
)

type monitorService struct {
	IMonitorRepository repository.IMonitorRepository
}

type IMonitorService interface {
	GetAll(page int, pageSize int) ([] model.Monitor)
	GetById(id uuid.UUID) (model.Monitor, error)
	SearchByName(page int, pageSize int, name string) ([]model.Monitor, error)
	Create(monitor model.Monitor) error
	Update(monitorDTO dto.MonitorDTO) error
	Delete(id uuid.UUID) error
}

func NewMonitorService(monitorRepository repository.IMonitorRepository) IMonitorService {
	return &monitorService{IMonitorRepository: monitorRepository}
}

func (monitorService *monitorService) GetAll(page int, pageSize int) []model.Monitor {
	return monitorService.IMonitorRepository.GetAll(page, pageSize)
}

func (monitorService *monitorService) GetById(id uuid.UUID) (model.Monitor, error) {
	return monitorService.IMonitorRepository.GetById(id)
}

func (monitorService *monitorService) SearchByName(page int, pageSize int, name string) ([]model.Monitor, error) {
	return monitorService.IMonitorRepository.SearchByName(page, pageSize, name)
}

func (monitorService *monitorService) Create(monitor model.Monitor) error {
	monitor.Product.Id = uuid.New()
	monitor.ProductId = monitor.Product.Id
	monitor.Product.Type = model.MONITOR
	return monitorService.IMonitorRepository.Create(monitor)
}

func (monitorService *monitorService) Update(monitorDTO dto.MonitorDTO) error {
	monitor, err := monitorService.GetById(monitorDTO.Product.Id)
	if err != nil {
		return err
	}
	updatedMonitor := mapper.ToMonitor(monitorDTO)
	updatedMonitor.Product.Id = monitor.Product.Id
	updatedMonitor.ProductId = monitor.Product.Id
	return monitorService.IMonitorRepository.Update(updatedMonitor)
}

func (monitorService *monitorService) Delete(id uuid.UUID) error {
	monitor, err := monitorService.GetById(id)
	if err != nil {
		return err
	}
	return monitorService.IMonitorRepository.Delete(monitor)
}