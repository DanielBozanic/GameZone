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
	GetAll() ([] model.Monitor)
	GetById(id uuid.UUID) (model.Monitor, error)
	GetByName(name string) (model.Monitor, error)
	Create(monitor model.Monitor) error
	Update(monitorDTO dto.MonitorDTO) error
	Delete(id uuid.UUID) error
}

func NewMonitorService(monitorRepository repository.IMonitorRepository) IMonitorService {
	return &monitorService{IMonitorRepository: monitorRepository}
}

func (monitorService *monitorService) GetAll() []model.Monitor {
	return monitorService.IMonitorRepository.GetAll()
}

func (monitorService *monitorService) GetById(id uuid.UUID) (model.Monitor, error) {
	return monitorService.IMonitorRepository.GetById(id)
}

func (monitorService *monitorService) GetByName(name string) (model.Monitor, error) {
	return monitorService.IMonitorRepository.GetByName(name)
}

func (monitorService *monitorService) Create(monitor model.Monitor) error {
	monitor.Id = uuid.New()
	return monitorService.IMonitorRepository.Create(monitor)
}

func (monitorService *monitorService) Update(monitorDTO dto.MonitorDTO) error {
	monitor, err := monitorService.GetById(monitorDTO.Id)
	if err != nil {
		return err
	}
	updatedMonitor := mapper.ToMonitor(monitorDTO)
	updatedMonitor.Id = monitor.Id
	return monitorService.IMonitorRepository.Update(updatedMonitor)
}

func (monitorService *monitorService) Delete(id uuid.UUID) error {
	monitor, err := monitorService.GetById(id)
	if err != nil {
		return err
	}
	return monitorService.IMonitorRepository.Delete(monitor)
}