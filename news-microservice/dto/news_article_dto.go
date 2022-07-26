package dto

import "time"

type NewsArticleDTO struct {
	Id          int
	UnpublishedTitle       string
	UnpublishedDescription string
	UnpublishedContent  string
	PublishedTitle       string
	PublishedDescription string 
	PublishedContent    string    
	DateTime    time.Time
	IsSent      *bool
	Archived    *bool
}