package api

import (
	"contact-and-report/dto"
	"contact-and-report/mapper"
	"contact-and-report/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportAPI struct {
	IReportService service.IReportService
}

func NewReportAPI(reportService service.IReportService) ReportAPI {
	return ReportAPI{IReportService: reportService}
}

func (reportApi *ReportAPI) GetReportsByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	reports := reportApi.IReportService.GetReportsByUserId(userId)
	c.JSON(http.StatusOK, mapper.ToReportDTOs(reports))
}

func (reportApi *ReportAPI) AddReport(c *gin.Context) {
	var reportDTO dto.ReportDTO
	err := c.BindJSON(&reportDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	report := mapper.ToReport(reportDTO)
	msg := reportApi.IReportService.AddReport(report)
	if msg == "" {
		c.JSON(http.StatusOK, "Report sent successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}