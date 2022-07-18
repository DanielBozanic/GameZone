package api

import (
	"product/dto"
	"product/dto/filter"
	"product/mapper"
	"product/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)


type HardDiskDriveAPI struct {
	IHardDiskDriveService service.IHardDiskDriveService
}

func NewHardDiskDriveAPI(hardDiskDriveService service.IHardDiskDriveService) HardDiskDriveAPI {
	return HardDiskDriveAPI{IHardDiskDriveService: hardDiskDriveService}
}

func (hardDiskDriveApi *HardDiskDriveAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	hardDiskDrives := hardDiskDriveApi.IHardDiskDriveService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToHardDiskDriveDTOs(hardDiskDrives))
}

func (hardDiskDriveApi *HardDiskDriveAPI) GetNumberOfRecords(c *gin.Context) {
	numberOfRecords := hardDiskDriveApi.IHardDiskDriveService.GetNumberOfRecords()
	c.JSON(http.StatusOK, numberOfRecords)
}

func (hardDiskDriveApi *HardDiskDriveAPI) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	hardDiskDrive, err := hardDiskDriveApi.IHardDiskDriveService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToHardDiskDriveDTO(hardDiskDrive))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (hardDiskDriveApi *HardDiskDriveAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	hardDiskDrives, err := hardDiskDriveApi.IHardDiskDriveService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToHardDiskDriveDTOs(hardDiskDrives))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (hardDiskDriveApi *HardDiskDriveAPI) GetNumberOfRecordsSearch(c *gin.Context) {
	numberOfRecords := hardDiskDriveApi.IHardDiskDriveService.GetNumberOfRecordsSearch(c.Query("name"))
	c.JSON(http.StatusOK, numberOfRecords)
}

func (hardDiskDriveApi *HardDiskDriveAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.HardDiskDriveFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hardDiskDrives, err := hardDiskDriveApi.IHardDiskDriveService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToHardDiskDriveDTOs(hardDiskDrives))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (hardDiskDriveApi *HardDiskDriveAPI) GetNumberOfRecordsFilter(c *gin.Context) {
	var filter filter.HardDiskDriveFilter
	err := c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	numberOfRecords := hardDiskDriveApi.IHardDiskDriveService.GetNumberOfRecordsFilter(filter)
	c.JSON(http.StatusOK, numberOfRecords)
}

func (hardDiskDriveApi *HardDiskDriveAPI) GetCapacities(c *gin.Context) {
	capacities := hardDiskDriveApi.IHardDiskDriveService.GetCapacities()
	c.JSON(http.StatusOK, capacities)
}

func (hardDiskDriveApi *HardDiskDriveAPI) GetForms(c *gin.Context) {
	forms := hardDiskDriveApi.IHardDiskDriveService.GetForms()
	c.JSON(http.StatusOK, forms)
}

func (hardDiskDriveApi *HardDiskDriveAPI) GetManufacturers(c *gin.Context) {
	manufacturers := hardDiskDriveApi.IHardDiskDriveService.GetManufacturers()
	c.JSON(http.StatusOK, manufacturers)
}

func (hardDiskDriveApi *HardDiskDriveAPI) GetDiskSpeeds(c *gin.Context) {
	diskSpeeds := hardDiskDriveApi.IHardDiskDriveService.GetDiskSpeeds()
	c.JSON(http.StatusOK, diskSpeeds)
}

func (hardDiskDriveApi *HardDiskDriveAPI) Create(c *gin.Context) {
	var hardDiskDriveDTO dto.HardDiskDriveDTO
	err := c.BindJSON(&hardDiskDriveDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hardDiskDrive := mapper.ToHardDiskDrive(hardDiskDriveDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := hardDiskDriveApi.IHardDiskDriveService.Create(hardDiskDrive)

	if msg == "" {
		c.JSON(http.StatusOK, "Hard disk drive added successfully.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (hardDiskDriveApi *HardDiskDriveAPI) Update(c *gin.Context) {
	var hardDiskDriveDTO dto.HardDiskDriveDTO
	err := c.BindJSON(&hardDiskDriveDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := hardDiskDriveApi.IHardDiskDriveService.Update(hardDiskDriveDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "Hard disk drive updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (hardDiskDriveApi *HardDiskDriveAPI) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := hardDiskDriveApi.IHardDiskDriveService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "Hard disk drive deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}