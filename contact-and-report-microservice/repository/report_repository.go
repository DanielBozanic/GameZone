package repository

import (
	"contact-and-report/model"

	"gorm.io/gorm"
)

type reportRepository struct {
	Database *gorm.DB
}

type IReportRepository interface {
	GetById(id int) (model.Report, error) 
	GetAll(page int, pageSize int) ([] model.Report)
	GetNumberOfRecords() int64
	GetUnansweredReportsByUserId(page int, pageSize int, userId int) []model.Report
	GetNumberOfRecordsUnansweredReportsByUserId(userId int) int64
	GetAnsweredReportsByUserId(page int, pageSize int, userId int) []model.Report
	GetNumberOfRecordsAnsweredReportsByUserId(userId int) int64
	Create(report model.Report) error
	Update(report model.Report) error
}

func NewReportRepository(DB *gorm.DB) IReportRepository {
	return &reportRepository{Database: DB}
}

func (reportRepo *reportRepository) GetById(id int) (model.Report, error) {
	var report model.Report
	result := reportRepo.Database.
		Where("id = ?", id).
		First(&report)
	return report, result.Error
}

func (reportRepo *reportRepository) GetAll(page int, pageSize int) []model.Report {
	var reports []model.Report
	offset := (page - 1) * pageSize
	reportRepo.Database.
		Offset(offset).Limit(pageSize).
		Find(&reports)
	return reports
}

func (reportRepo *reportRepository) GetNumberOfRecords() int64 {
	var count int64
	reportRepo.Database.
		Model(&model.Report{}).
		Count(&count)
	return count
}

func (reportRepo *reportRepository) GetUnansweredReportsByUserId(page int, pageSize int, userId int) []model.Report {
	var reports []model.Report
	offset := (page - 1) * pageSize
	reportRepo.Database.
		Offset(offset).Limit(pageSize).
		Where("user_id = ? AND answered = false", userId).
		Order("date_time DESC").
		Find(&reports)
	return reports
}

func (reportRepo *reportRepository) GetNumberOfRecordsUnansweredReportsByUserId(userId int) int64 {
	var count int64
	reportRepo.Database.
		Model(&model.Report{}).
		Where("user_id = ? AND answered = false", userId).
		Count(&count)
	return count
}

func (reportRepo *reportRepository) GetAnsweredReportsByUserId(page int, pageSize int, userId int) []model.Report {
	var reports []model.Report
	offset := (page - 1) * pageSize
	reportRepo.Database.
		Offset(offset).Limit(pageSize).
		Where("user_id = ? AND answered = true", userId).
		Order("date_time DESC").
		Find(&reports)
	return reports
}

func (reportRepo *reportRepository) GetNumberOfRecordsAnsweredReportsByUserId(userId int) int64 {
	var count int64
	reportRepo.Database.
		Model(&model.Report{}).
		Where("user_id = ? AND answered = true", userId).
		Count(&count)
	return count
}

func (reportRepo *reportRepository) Create(report model.Report) error {
	result := reportRepo.Database.Create(&report)
	return result.Error
}

func (reportRepo *reportRepository) Update(report model.Report) error {
	result := reportRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&report)
	return result.Error
}