package service

import (
	"bytes"
	"contact-and-report/model"
	"contact-and-report/repository"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type banService struct {
	IBanRepository repository.IBanRepository
	IReportRepository repository.IReportRepository
}

type IBanService interface {
	GetUserBanHistory(userId int) []model.Ban
	IsUserBanned(userId int) bool
	AddBan(ban model.Ban, reportId int) string
	SendEmailToBannedUser(ban model.Ban) string
}

func NewBanService(banRepository repository.IBanRepository) IBanService {
	return &banService{IBanRepository: banRepository}
}

func (banService *banService) GetUserBanHistory(userId int) []model.Ban {
	return banService.IBanRepository.GetUserBanHistory(userId)
}

func (banService *banService) IsUserBanned(userId int) bool {
	_, err := banService.IBanRepository.IsUserBanned(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (banService *banService) AddBan(ban model.Ban, reportId int) string {
	_, err := banService.IBanRepository.IsUserBanned(ban.UserId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "User is already banned"
	}
	
	err = banService.IBanRepository.Create(ban)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (banService *banService) SendEmailToBannedUser(ban model.Ban) string {
	req, err := http.NewRequest("GET", "http://localhost:5000/api/users/getById?userId=" +  strconv.Itoa(ban.UserId), nil)
	client := &http.Client{}
	resp, err := client.Do(req)

	recipients := []string{}
	var target map[string]interface{}
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&target)
	email := target["user"].(map[string]interface{})["email"].(string)
	recipients = append(recipients, email)

	data := map[string]interface{}{
		"subject": "You have been banned" ,
		"recipients": recipients,
		"content": map[string]interface{}{
			"template": "banned_user",
			"params": map[string]interface{}{
				"reason": ban.Reason.String(),
				"expirationDate": ban.ExpirationDate.Format("02-Jan-2006 15:04:05"),
			},
		},
	}
	jsonData, _ := json.Marshal(data)

	req, err = http.NewRequest("POST", "http://localhost:5001/api/email/sendEmail", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&target)

	return ""
}