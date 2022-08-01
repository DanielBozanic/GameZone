package dto

import (
	"time"
)

type ReportDTO struct {
	Id                int
	UserId            int
	Reason            string
	Description       string
	DateTime          time.Time
}