//go:build wireinject
// +build wireinject

package di

import (
	"comment-and-rating/api"
	"comment-and-rating/repository"
	"comment-and-rating/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitProductCommentAPI(db *gorm.DB) api.ProductCommentAPI {
	wire.Build(repository.NewProductCommentRepository, service.NewProductCommentService, api.NewProductCommentAPI)
	return api.ProductCommentAPI{}
}