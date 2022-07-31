package repository

import (
	"contact-and-report/model"

	"gorm.io/gorm"
)

type reportRepository struct {
	Database *gorm.DB
}

type IReportRepository interface {
	GetReportsByUserId(userId int) []model.Report
	Create(report model.Report) error
}

func NewReportRepository(DB *gorm.DB) IReportRepository {
	return &reportRepository{Database: DB}
}

func (reportRepo *reportRepository) GetReportsByUserId(userId int) []model.Report {
	var reports []model.Report
	reportRepo.Database.
		Where("user_id = ?", userId).
		Order("date_time DESC").
		Find(&reports)
	return reports
}

func (reportRepo *reportRepository) Create(report model.Report) error {
	result := reportRepo.Database.Create(&report)
	return result.Error
}