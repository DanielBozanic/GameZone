package dto

import (
	"comment-and-rating/model"
	"time"
)

type ProductCommentDTO struct {
	Id        int
	ProductName string
	Username  string
	Comment   string
	Rating    model.Rating
	Archived  *bool
	DateTime time.Time
}