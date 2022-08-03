package repository

import (
	"contact-and-report/model"

	"gorm.io/gorm"
)

type contactRepository struct {
	Database *gorm.DB
}

type IContactRepository interface {
	GetById(id int) (model.ContactMessage, error)
	GetUnansweredContactMessages() []model.ContactMessage
	GetUnansweredContactMessagesByUserId(userId int) []model.ContactMessage
	GetContactMessagesByUserId(userId int) []model.ContactMessage
	Create(contactMsg model.ContactMessage) error
	Update(contactMsg model.ContactMessage) error
}

func NewContactRepository(DB *gorm.DB) IContactRepository {
	return &contactRepository{Database: DB}
}

func (contactRepo *contactRepository) GetById(id int) (model.ContactMessage, error) {
	var contactMessage model.ContactMessage
	result := contactRepo.Database.
		Where("id = ?", id).
		First(&contactMessage)
	return contactMessage, result.Error
}

func (contactRepo *contactRepository) GetUnansweredContactMessages() []model.ContactMessage {
	var contactMessages []model.ContactMessage
	contactRepo.Database.
		Where("answer IS NULL").
		Order("date_time DESC").
		Find(&contactMessages)
	return contactMessages
}

func (contactRepo *contactRepository) GetUnansweredContactMessagesByUserId(userId int) []model.ContactMessage {
	var contactMessages []model.ContactMessage
	contactRepo.Database.
		Where("user_id = ? AND answer IS NULL", userId).
		Order("date_time DESC").
		Find(&contactMessages)
	return contactMessages
}

func (contactRepo *contactRepository) GetContactMessagesByUserId(userId int) []model.ContactMessage {
	var contactMessages []model.ContactMessage
	contactRepo.Database.
		Where("user_id = ?", userId).
		Order("date_time DESC").
		Find(&contactMessages)
	return contactMessages
}

func (contactRepo *contactRepository) Create(contactMsg model.ContactMessage) error {
	result := contactRepo.Database.Create(&contactMsg)
	return result.Error
}

func (contactRepo *contactRepository) Update(contactMsg model.ContactMessage) error {
	result := contactRepo.Database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&contactMsg)
	return result.Error
}