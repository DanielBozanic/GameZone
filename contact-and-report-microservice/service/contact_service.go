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
	GetUnansweredContactMessagesByUserId(userId int) []model.ContactMessage
	GetAnsweredContactMessagesByUserId(userId int) []model.ContactMessage
	SendContactMessage(contactMsg model.ContactMessage, userId int) string
	AnswerContactMessage(contactMsgDTO dto.ContactMessageDTO) string
}

func NewContactService(contactRepository repository.IContactRepository) IContactService {
	return &contactService{IContactRepository: contactRepository}
}

func (contactService *contactService) GetUnansweredContactMessagesByUserId(userId int) []model.ContactMessage {
	return contactService.IContactRepository.GetUnansweredContactMessagesByUserId(userId)
}

func (contactService *contactService) GetAnsweredContactMessagesByUserId(userId int) []model.ContactMessage {
	return contactService.IContactRepository.GetAnsweredContactMessagesByUserId(userId)
}

func (contactService *contactService) SendContactMessage(contactMsg model.ContactMessage, userId int) string {
	contactMsg.UserId = userId
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