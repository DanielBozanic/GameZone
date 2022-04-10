package mapper

import (
	"product/dto"
	"product/model"
)


func ToMonitor(monitorDTO dto.MonitorDTO) (model.Monitor) {
	return model.Monitor {
		Name: monitorDTO.Name, 
		Price: monitorDTO.Price, 
		Size: monitorDTO.Size,
		AspectRatio: monitorDTO.AspectRatio,
		Resolution: monitorDTO.Resolution,
		ContrastRatio: monitorDTO.ContrastRatio,
		ResponseTime: monitorDTO.ResponseTime,
		PanelType: monitorDTO.PanelType,
		ViewingAngle: monitorDTO.ViewingAngle,
		Brightness: monitorDTO.Brightness,
		RefreshRate: monitorDTO.RefreshRate,
		Amount: monitorDTO.Amount,
		Manufacturer: monitorDTO.Manufacturer,
	}
}

func ToMonitorDTO(monitor model.Monitor) dto.MonitorDTO {
	return dto.MonitorDTO {
		Id: monitor.Id, 
		Name: monitor.Name, 
		Price: monitor.Price, 
		Size: monitor.Size,
		AspectRatio: monitor.AspectRatio,
		Resolution: monitor.Resolution,
		ContrastRatio: monitor.ContrastRatio,
		ResponseTime: monitor.ResponseTime,
		PanelType: monitor.PanelType,
		ViewingAngle: monitor.ViewingAngle,
		Brightness: monitor.Brightness,
		RefreshRate: monitor.RefreshRate,
		Amount: monitor.Amount,
		Manufacturer: monitor.Manufacturer,
	}
}

func ToMonitorDTOs(monitors []model.Monitor) []dto.MonitorDTO {
	monitorDTOs := make([]dto.MonitorDTO, len(monitors))

	for i, itm := range monitors {
		monitorDTOs[i] = ToMonitorDTO(itm)
	}

	return monitorDTOs
}