//go:build wireinject
// +build wireinject

package di

import (
	"news/api"
	"news/repository"
	"news/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitNewsArticleAPI(db *gorm.DB) api.NewsArticleAPI {
	wire.Build(repository.NewNewsArticleRepository, service.NewNewsArticleService, api.NewNewsArticleAPI)
	return api.NewsArticleAPI{}
}

func InitNewsCommentAPI(db *gorm.DB) api.NewsCommentAPI {
	wire.Build(repository.NewNewsCommentRepository, service.NewNewsCommentService, api.NewNewsCommentAPI)
	return api.NewsCommentAPI{}
}

func InitNewsSubscriptionAPI(db *gorm.DB) api.NewsSubscriptionAPI {
	wire.Build(repository.NewNewsSubscriptionRepository, repository.NewNewsArticleRepository, service.NewNewsSubscriptionService, api.NewNewsSubscriptionAPI)
	return api.NewsSubscriptionAPI{}
}