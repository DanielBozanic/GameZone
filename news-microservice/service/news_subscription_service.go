package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"news/model"
	"news/repository"
	"strconv"

	"gorm.io/gorm"
)

type newsSubscriptionService struct {
	INewsSubscriptionRepository repository.INewsSubscriptionRepository
	INewsArticleRepository repository.INewsArticleRepository
}

type INewsSubscriptionService interface {
	Subscribe(userId int) string
	Unsubscribe(userId int) string
	IsUserSubscribed(userId int) bool
	SendEmails() string
}

func NewNewsSubscriptionService(
	newsSubscriptionRepository repository.INewsSubscriptionRepository, 
	newsArticleRepository repository.INewsArticleRepository) INewsSubscriptionService {
	return &newsSubscriptionService{
		INewsSubscriptionRepository: newsSubscriptionRepository, 
		INewsArticleRepository: newsArticleRepository}
}

func (newsSubscriptionService *newsSubscriptionService) Subscribe(userId int) string {
	msg := ""
	var newsSubscription model.NewsSubscription
	_, err := newsSubscriptionService.INewsSubscriptionRepository.IsUserSubscribed(userId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "You are already subscribed to receive emails."
	}
	newsSubscription.UserId = userId
	err = newsSubscriptionService.INewsSubscriptionRepository.Create(newsSubscription)
	if err != nil {
		msg = err.Error()
	}
	return msg
}

func (newsSubscriptionService *newsSubscriptionService) Unsubscribe(userId int) string {
	msg := ""
	newsSubscription, err := newsSubscriptionService.INewsSubscriptionRepository.IsUserSubscribed(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "You are not subscribed to receive emails."
	}
	
	err = newsSubscriptionService.INewsSubscriptionRepository.Delete(newsSubscription)
	if err != nil {
		msg = err.Error()
	}
	return msg
}

func (newsSubscriptionService *newsSubscriptionService) IsUserSubscribed(userId int) bool {
	_, err := newsSubscriptionService.INewsSubscriptionRepository.IsUserSubscribed(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	} else {
		return true
	}
}

func (newsSubscriptionService *newsSubscriptionService) SendEmails() string {
	newsSubscriptions := newsSubscriptionService.INewsSubscriptionRepository.GetAll()
	unsentNewsArticles := newsSubscriptionService.INewsArticleRepository.GetUnsentPublishedArticles()
	recipients := []string{} 
	for _, newsSubscription := range newsSubscriptions {
		req, err := http.NewRequest("GET", "http://localhost:5000/api/users/getById?userId=" +  strconv.Itoa(newsSubscription.UserId), nil)
		client := &http.Client{}
		resp, err := client.Do(req)

		var target map[string]interface{}
		if err != nil {
			return err.Error()
		}
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&target)
		email := target["user"].(map[string]interface{})["email"].(string)
		recipients = append(recipients, email)
	}

	if len(recipients) > 0 {
		newsArticlesForSending := ""
		for _, unsentNewsArticle := range unsentNewsArticles {
			newsArticlesForSending += `<h2 style="text-align: center;">` + 
				unsentNewsArticle.PublishedTitle + `</h2>` + 
				unsentNewsArticle.PublishedContent + `<br/><hr/>`
		}

		data := map[string]interface{}{
			"subject": "Weekly news articles",
			"recipients": recipients,
			"content": map[string]interface{}{
				"template": "news_subscription",
				"params": map[string]interface{}{
					"htmlNews": newsArticlesForSending,
				},
			},
		}
		jsonData, _ := json.Marshal(data)

		req, err := http.NewRequest("POST", "http://localhost:5001/api/email/sendEmail", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err.Error()
		}

		var target interface{}
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(target)

		for _, unsentNewsArticle := range unsentNewsArticles {
			*unsentNewsArticle.IsSent = true
			newsSubscriptionService.INewsArticleRepository.Update(unsentNewsArticle)
		}
	}

	return ""
}