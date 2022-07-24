package dto

import "time"

type NewsCommentDTO struct {
	Id          int
	NewsArticle NewsArticleDTO
	Comment     string
	Username    string
	DateTime    time.Time
	Archived    *bool
}