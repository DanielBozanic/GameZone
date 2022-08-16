package service

import (
	"contact-and-report/dto"
	"contact-and-report/model"
	"contact-and-report/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type contactService struct {
	IContactRepository repository.IContactRepository
}

type IContactService interface {
	GetContactMessages() []model.ContactMessage 
	GetContactMessagesByUserId(userId int) []model.ContactMessage
	GetNumberOfUnansweredContactMessagesByUserId(userId int) int64
	SendContactMessage(contactMsg model.ContactMessage, userData dto.UserData) string
	AnswerContactMessage(contactMsgDTO dto.ContactMessageDTO) string
}

func NewContactService(contactRepository repository.IContactRepository) IContactService {
	return &contactService{IContactRepository: contactRepository}
}

func (contactService *contactService) GetContactMessages() []model.ContactMessage {
	return contactService.IContactRepository.GetContactMessages()
}

func (contactService *contactService) GetContactMessagesByUserId(userId int) []model.ContactMessage {
	return contactService.IContactRepository.GetContactMessagesByUserId(userId)
}

func (contactService *contactService) GetNumberOfUnansweredContactMessagesByUserId(userId int) int64 {
	return contactService.IContactRepository.GetNumberOfUnansweredContactMessagesByUserId(userId)
}

func (contactService *contactService) SendContactMessage(contactMsg model.ContactMessage, userData dto.UserData) string {
	contactMsg.UserId = userData.Id
	contactMsg.Username = userData.Username
	contactMsg.DateTime = time.Now()
	err := contactService.IContactRepository.Create(contactMsg)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (contactService *contactService) AnswerContactMessage(contactMsgDTO dto.ContactMessageDTO) string {
	contactMsg, err := contactService.IContactRepository.GetById(contactMsgDTO.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err.Error()
	}

	contactMsg.Answer = contactMsgDTO.Answer
	err = contactService.IContactRepository.Update(contactMsg)
	if err != nil {
		return err.Error()
	}
	return ""
}