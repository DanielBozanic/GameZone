package dto

import "time"

type NewsArticleDTO struct {
	Id          int
	Title       string
	Description *string
	UnpublishedContent  string    
	PublishedContent    *string    
	DateTime    time.Time
	IsSent      *bool
	Archived    *bool
}