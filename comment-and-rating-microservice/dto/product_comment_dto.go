package dto

import (
	"comment-and-rating/model"
	"time"
)

type ProductCommentDTO struct {
	Id        int
	ProductId int
	UserId    int
	Username  string
	Comment   string
	Rating    model.Rating
	Archived  *bool
	DateTime time.Time
}