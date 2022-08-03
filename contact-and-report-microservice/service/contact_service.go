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
	GetUnansweredContactMessages() []model.ContactMessage 
	GetUnansweredContactMessagesByUserId(userId int) []model.ContactMessage
	GetContactMessagesByUserId(userId int) []model.ContactMessage
	SendContactMessage(contactMsg model.ContactMessage, userData dto.UserData) string
	AnswerContactMessage(contactMsgDTO dto.ContactMessageDTO) string
}

func NewContactService(contactRepository repository.IContactRepository) IContactService {
	return &contactService{IContactRepository: contactRepository}
}

func (contactService *contactService) GetUnansweredContactMessages() []model.ContactMessage {
	return contactService.IContactRepository.GetUnansweredContactMessages()
}

func (contactService *contactService) GetUnansweredContactMessagesByUserId(userId int) []model.ContactMessage {
	return contactService.IContactRepository.GetUnansweredContactMessagesByUserId(userId)
}

func (contactService *contactService) GetContactMessagesByUserId(userId int) []model.ContactMessage {
	return contactService.IContactRepository.GetContactMessagesByUserId(userId)
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