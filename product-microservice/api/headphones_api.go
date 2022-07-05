package api

import (
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type HeadphonesAPI struct {
	IHeadphonesService service.IHeadphonesService
}

func NewHeadphonesAPI(headphonesService service.IHeadphonesService) HeadphonesAPI {
	return HeadphonesAPI{IHeadphonesService: headphonesService}
}

func (headphonesApi *HeadphonesAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	headphoness := headphonesApi.IHeadphonesService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToHeadphonesDTOs(headphoness))
}

func (headphonesApi *HeadphonesAPI) GetNumberOfRecords(c *gin.Context) {
	numberOfRecords := headphonesApi.IHeadphonesService.GetNumberOfRecords()
	c.JSON(http.StatusOK, numberOfRecords)
}

func (headphonesApi *HeadphonesAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	headphones, err := headphonesApi.IHeadphonesService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToHeadphonesDTO(headphones))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (headphonesApi *HeadphonesAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	headphones, err := headphonesApi.IHeadphonesService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToHeadphonesDTOs(headphones))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (headphonesApi *HeadphonesAPI) GetNumberOfRecordsSearch(c *gin.Context) {
	numberOfRecords := headphonesApi.IHeadphonesService.GetNumberOfRecordsSearch(c.Query("name"))
	c.JSON(http.StatusOK, numberOfRecords)
}

func (headphonesApi *HeadphonesAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.HeadphonesFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	headphones, err := headphonesApi.IHeadphonesService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToHeadphonesDTOs(headphones))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (headphonesApi *HeadphonesAPI) GetNumberOfRecordsFilter(c *gin.Context) {
	var filter filter.HeadphonesFilter
	err := c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	numberOfRecords := headphonesApi.IHeadphonesService.GetNumberOfRecordsFilter(filter)
	c.JSON(http.StatusOK, numberOfRecords)
}

func (headphonesApi *HeadphonesAPI) GetManufacturers(c *gin.Context) {
	manufacturers := headphonesApi.IHeadphonesService.GetManufacturers()
	c.JSON(http.StatusOK, manufacturers)
}

func (headphonesApi *HeadphonesAPI) GetConnectionTypes(c *gin.Context) {
	connectionTypes := headphonesApi.IHeadphonesService.GetConnectionTypes()
	c.JSON(http.StatusOK, connectionTypes)
}

func (headphonesApi *HeadphonesAPI) Create(c *gin.Context) {
	var headphonesDTO dto.HeadphonesDTO
	err := c.BindJSON(&headphonesDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	headphones := mapper.ToHeadphones(headphonesDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := headphonesApi.IHeadphonesService.Create(headphones)

	if error == nil {
		c.JSON(http.StatusOK, "Headphones stored successfully.")
	} else {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}

func (headphonesApi *HeadphonesAPI) Update(c *gin.Context) {
	var headphonesDTO dto.HeadphonesDTO
	err := c.BindJSON(&headphonesDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := headphonesApi.IHeadphonesService.Update(headphonesDTO)

	if error == nil {
		c.JSON(http.StatusOK, "Headphones updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	} 
}

func (headphonesApi *HeadphonesAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := headphonesApi.IHeadphonesService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "Headphones deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}