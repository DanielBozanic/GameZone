package api

import (
	"news/middleware"
	"news/service"

	"net/http"

	"github.com/gin-gonic/gin"
)


type NewsSubscriptionAPI struct {
	INewsSubscriptionService service.INewsSubscriptionService
}

func NewNewsSubscriptionAPI(newsSubscriptionService service.INewsSubscriptionService) NewsSubscriptionAPI {
	return NewsSubscriptionAPI{INewsSubscriptionService: newsSubscriptionService}
}

func (newsSubscriptionApi *NewsSubscriptionAPI) Subscribe(c *gin.Context) {
	userData := middleware.GetUserData(c)
	msg := newsSubscriptionApi.INewsSubscriptionService.Subscribe(userData.Id)

	if msg == "" {
		c.JSON(http.StatusOK, "You are subscribed.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (newsSubscriptionApi *NewsSubscriptionAPI) Unsubscribe(c *gin.Context) {
	userData := middleware.GetUserData(c)
	msg := newsSubscriptionApi.INewsSubscriptionService.Unsubscribe(userData.Id)

	if msg == "" {
		c.JSON(http.StatusOK, "You are unsubscribed.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (newsSubscriptionApi *NewsSubscriptionAPI) IsUserSubscribed(c *gin.Context) {
	userData := middleware.GetUserData(c)
	isUserSubscribed := newsSubscriptionApi.INewsSubscriptionService.IsUserSubscribed(userData.Id)
	c.JSON(http.StatusOK, isUserSubscribed)
}