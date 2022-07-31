package dto

import (
	"contact-and-report/model"
	"time"
)

type ReportDTO struct {
	Id                int
	UserId            int
	Reason            model.Reason
	ReasonDescription string
	DateTime          time.Time
}