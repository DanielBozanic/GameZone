package dto

import "time"

type ContactMessageDTO struct {
	Id           int
	UserId       int
	UserQuestion string
	Answer       string
	DateTime     time.Time
}