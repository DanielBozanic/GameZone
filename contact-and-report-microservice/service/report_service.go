package service

import (
	"contact-and-report/model"
	"contact-and-report/repository"
	"time"
)

type reportService struct {
	IReportRepository repository.IReportRepository
}

type IReportService interface {
	GetReportsByUserId(userId int) []model.Report
	AddReport(report model.Report) string
}

func NewReportService(reportRepository repository.IReportRepository) IReportService {
	return &reportService{IReportRepository: reportRepository}
}

func (reportService *reportService) GetReportsByUserId(userId int) []model.Report {
	return reportService.IReportRepository.GetReportsByUserId(userId)
}

func (reportService *reportService) AddReport(report model.Report) string {
	report.DateTime = time.Now()
	err := reportService.IReportRepository.Create(report)
	if err != nil {
		return err.Error()
	}
	return ""
}