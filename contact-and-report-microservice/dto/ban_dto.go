package dto

import (
	"contact-and-report/model"
	"time"
)

type BanDTO struct {
	Id             int
	UserId         int
	Reason         model.Reason
	ExpirationDate time.Time
}