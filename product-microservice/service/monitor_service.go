package service

import (
	"errors"
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/model"
	"product/repository"

	"github.com/go-sql-driver/mysql"
)

type monitorService struct {
	IMonitorRepository repository.IMonitorRepository
}

type IMonitorService interface {
	GetAll(page int, pageSize int) ([] model.Monitor)
	GetNumberOfRecords() int64
	GetById(id int) (model.Monitor, error)
	SearchByName(page int, pageSize int, name string) ([]model.Monitor, error)
	GetNumberOfRecordsSearch(name string) int64
	Filter(page int, pageSize int, filter filter.MonitorFilter) ([]model.Monitor, error)
	GetNumberOfRecordsFilter(filter filter.MonitorFilter) int64
	GetManufacturers() []string
	GetAspectRatios() []string
	GetResolutions() []string
	GetRefreshRates() []string
	Create(monitor model.Monitor) string
	Update(monitorDTO dto.MonitorDTO) string
	Delete(id int) error
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

func (monitorService *monitorService) GetById(id int) (model.Monitor, error) {
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

func (monitorService *monitorService) Create(monitor model.Monitor) string {
	msg := ""
	monitor.Product.Type = model.MONITOR
	err := monitorService.IMonitorRepository.Create(monitor)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (monitorService *monitorService) Update(monitorDTO dto.MonitorDTO) string {
	msg := ""
	monitor, err := monitorService.GetById(monitorDTO.Product.Id)
	if err != nil {
		return err.Error()
	}
	updatedMonitor := mapper.ToMonitor(monitorDTO)
	updatedMonitor.Product.Id = monitor.Product.Id
	updatedMonitor.ProductId = monitor.Product.Id
	updatedMonitor.Product.Image.Id = monitor.Product.Image.Id
	updatedMonitor.Product.ImageId = monitor.Product.Image.Id
	err = monitorService.IMonitorRepository.Update(updatedMonitor)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		msg = "Product with this name already exists"
	}
	return msg
}

func (monitorService *monitorService) Delete(id int) error {
	monitor, err := monitorService.GetById(id)
	if err != nil {
		return err
	}
	monitor.Product.Archived = true
	return monitorService.IMonitorRepository.Update(monitor)
}