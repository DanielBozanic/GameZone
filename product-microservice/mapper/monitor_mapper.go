package mapper

import (
	"product/dto"
	"product/model"
)


func ToMonitor(monitorDTO dto.MonitorDTO) (model.Monitor) {
	return model.Monitor { 
		Product: ToProduct(monitorDTO.Product),
		Size: monitorDTO.Size,
		AspectRatio: monitorDTO.AspectRatio,
		Resolution: monitorDTO.Resolution,
		ContrastRatio: monitorDTO.ContrastRatio,
		ResponseTime: monitorDTO.ResponseTime,
		PanelType: monitorDTO.PanelType,
		ViewingAngle: monitorDTO.ViewingAngle,
		Brightness: monitorDTO.Brightness,
		RefreshRate: monitorDTO.RefreshRate,
	}
}

func ToMonitorDTO(monitor model.Monitor) dto.MonitorDTO {
	return dto.MonitorDTO { 
		Product: ToProductDTO(monitor.Product),
		Size: monitor.Size,
		AspectRatio: monitor.AspectRatio,
		Resolution: monitor.Resolution,
		ContrastRatio: monitor.ContrastRatio,
		ResponseTime: monitor.ResponseTime,
		PanelType: monitor.PanelType,
		ViewingAngle: monitor.ViewingAngle,
		Brightness: monitor.Brightness,
		RefreshRate: monitor.RefreshRate,
	}
}

func ToMonitorDTOs(monitors []model.Monitor) []dto.MonitorDTO {
	monitorDTOs := make([]dto.MonitorDTO, len(monitors))

	for i, itm := range monitors {
		monitorDTOs[i] = ToMonitorDTO(itm)
	}

	return monitorDTOs
}