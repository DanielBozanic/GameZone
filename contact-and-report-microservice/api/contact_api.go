package api

import (
	"contact-and-report/dto"
	"contact-and-report/mapper"
	"contact-and-report/middleware"
	"contact-and-report/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactAPI struct {
	IContactService service.IContactService
}

func NewContactAPI(contactService service.IContactService) ContactAPI {
	return ContactAPI{IContactService: contactService}
}

func (contactApi *ContactAPI) GetContactMessages(c *gin.Context) {
	contactMessages := contactApi.IContactService.GetContactMessages()
	c.JSON(http.StatusOK, mapper.ToContactMessageDTOs(contactMessages))
}

func (contactApi *ContactAPI) GetContactMessagesByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	contactMessages := contactApi.IContactService.GetContactMessagesByUserId(userId)
	c.JSON(http.StatusOK, mapper.ToContactMessageDTOs(contactMessages))
}

func (contactApi *ContactAPI) GetNumberOfUnansweredContactMessagesByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	numberOfUnansweredMsgs := contactApi.IContactService.GetNumberOfUnansweredContactMessagesByUserId(userId)
	c.JSON(http.StatusOK, numberOfUnansweredMsgs)
}

func (contactApi *ContactAPI) SendContactMessage(c *gin.Context) {
	var contactMsgDTO dto.ContactMessageDTO
	err := c.BindJSON(&contactMsgDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	contactMsg := mapper.ToContactMessage(contactMsgDTO)
	userData := middleware.GetUserData(c)
	msg := contactApi.IContactService.SendContactMessage(contactMsg, userData)
	if msg == "" {
		c.JSON(http.StatusOK, "Message sent successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (contactApi *ContactAPI) AnswerContactMessage(c *gin.Context) {
	var contactMsgDTO dto.ContactMessageDTO
	err := c.BindJSON(&contactMsgDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := contactApi.IContactService.AnswerContactMessage(contactMsgDTO)
	if msg == "" {
		c.JSON(http.StatusOK, "Message answered.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}