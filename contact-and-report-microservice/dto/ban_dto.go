package dto

import (
	"time"
)

type BanDTO struct {
	Id             int
	UserId         int
	Reason         string
	Description    string
	ExpirationDate time.Time
}