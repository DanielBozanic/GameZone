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
	GetUnansweredReportsByUserId(page int, pageSize int, userId int) []model.Report
	GetNumberOfRecordsUnansweredReportsByUserId(userId int) int64
	GetAnsweredReportsByUserId(page int, pageSize int, userId int) []model.Report
	GetNumberOfRecordsAnsweredReportsByUserId(userId int) int64
	AddReport(report model.Report) string
}

func NewReportService(reportRepository repository.IReportRepository) IReportService {
	return &reportService{IReportRepository: reportRepository}
}

func (reportService *reportService) GetUnansweredReportsByUserId(page int, pageSize int, userId int) []model.Report {
	return reportService.IReportRepository.GetUnansweredReportsByUserId(page, pageSize, userId)
}

func (reportService *reportService) GetNumberOfRecordsUnansweredReportsByUserId(userId int) int64 {
	return reportService.IReportRepository.GetNumberOfRecordsUnansweredReportsByUserId(userId)
}

func (reportService *reportService) GetAnsweredReportsByUserId(page int, pageSize int, userId int) []model.Report {
	return reportService.IReportRepository.GetAnsweredReportsByUserId(page, pageSize, userId)
}

func (reportService *reportService) GetNumberOfRecordsAnsweredReportsByUserId(userId int) int64 {
	return reportService.IReportRepository.GetNumberOfRecordsAnsweredReportsByUserId(userId)
}

func (reportService *reportService) AddReport(report model.Report) string {
	report.DateTime = time.Now()
	err := reportService.IReportRepository.Create(report)
	if err != nil {
		return err.Error()
	}
	return ""
}