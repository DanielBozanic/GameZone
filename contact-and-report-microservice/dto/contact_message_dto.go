package dto

import "time"

type ContactMessageDTO struct {
	Id           int
	UserId       int
	Username     string
	Message      string
	Answer       string
	DateTime     time.Time
}