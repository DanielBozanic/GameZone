package mapper

import (
	"contact-and-report/dto"
	"contact-and-report/model"
)


func ToReport(reportDTO dto.ReportDTO) (model.Report) {
	return model.Report {
		Id: reportDTO.Id,
		UserId: reportDTO.UserId,
		Reason: reportDTO.Reason,
		ReasonDescription: reportDTO.ReasonDescription,
		DateTime: reportDTO.DateTime,
		Answered: reportDTO.Answered,
	}
}

func ToReportDTO(report model.Report) dto.ReportDTO {
	return dto.ReportDTO {
		Id: report.Id,
		UserId: report.UserId,
		Reason: report.Reason,
		ReasonDescription: report.ReasonDescription,
		DateTime: report.DateTime,
		Answered: report.Answered,
	}
}

func ToReportDTOs(reports []model.Report) []dto.ReportDTO {
	reportDTOs := make([]dto.ReportDTO, len(reports))

	for i, itm := range reports {
		reportDTOs[i] = ToReportDTO(itm)
	}

	return reportDTOs
}