package service

import (
	"product/dto"
	"product/dto/filter"
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
	GetNumberOfRecords() int64
	GetById(id uuid.UUID) (model.Monitor, error)
	SearchByName(page int, pageSize int, name string) ([]model.Monitor, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.MonitorFilter) ([]model.Monitor, error)
	GetNumberOfRecordsFilter(filter filter.MonitorFilter) int64
	GetManufacturers() []string
	GetAspectRatios() []string
	GetResolutions() []string
	GetRefreshRates() []string
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

func (monitorService *monitorService) GetNumberOfRecords() int64 {
	return monitorService.IMonitorRepository.GetNumberOfRecords()
}

func (monitorService *monitorService) GetById(id uuid.UUID) (model.Monitor, error) {
	return monitorService.IMonitorRepository.GetById(id)
}

func (monitorService *monitorService) SearchByName(page int, pageSize int, name string) ([]model.Monitor, error) {
	return monitorService.IMonitorRepository.SearchByName(page, pageSize, name)
}

func (monitorService *monitorService) GetNumberOfRecordsSearch(name string) int64 {
	return monitorService.IMonitorRepository.GetNumberOfRecordsSearch(name)
}

func (monitorService *monitorService) Filter(page int, pageSize int, filter filter.MonitorFilter) ([]model.Monitor, error) {
	return monitorService.IMonitorRepository.Filter(page, pageSize, filter)
}

func (monitorService *monitorService) GetNumberOfRecordsFilter(filter filter.MonitorFilter) int64 {
	return monitorService.IMonitorRepository.GetNumberOfRecordsFilter(filter)
}

func (monitorService *monitorService) GetManufacturers() []string {
	return monitorService.IMonitorRepository.GetManufacturers()
}

func (monitorService *monitorService) GetAspectRatios() []string {
	return monitorService.IMonitorRepository.GetAspectRatios()
}

func (monitorService *monitorService) GetResolutions() []string {
	return monitorService.IMonitorRepository.GetResolutions()
}

func (monitorService *monitorService) GetRefreshRates() []string {
	return monitorService.IMonitorRepository.GetRefreshRates()
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