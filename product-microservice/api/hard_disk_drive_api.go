package api

import (
	"product/dto"
	"product/mapper"
	"product/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }
	
	hardDiskDrives := hardDiskDriveApi.IHardDiskDriveService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, gin.H{"hard_disk_drives": mapper.ToHardDiskDriveDTOs(hardDiskDrives)})
}

func (hardDiskDriveApi *HardDiskDriveAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	hardDiskDrive, err := hardDiskDriveApi.IHardDiskDriveService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"hard_disk_drive": mapper.ToHardDiskDriveDTO(hardDiskDrive)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (hardDiskDriveApi *HardDiskDriveAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

	hardDiskDrives, err := hardDiskDriveApi.IHardDiskDriveService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"hard_disk_drives": mapper.ToHardDiskDriveDTOs(hardDiskDrives)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (hardDiskDriveApi *HardDiskDriveAPI) Create(c *gin.Context) {
	var hardDiskDriveDTO dto.HardDiskDriveDTO
	err := c.BindJSON(&hardDiskDriveDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hardDiskDrive := mapper.ToHardDiskDrive(hardDiskDriveDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := hardDiskDriveApi.IHardDiskDriveService.Create(hardDiskDrive)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Hard disk drive stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (hardDiskDriveApi *HardDiskDriveAPI) Update(c *gin.Context) {
	var hardDiskDriveDTO dto.HardDiskDriveDTO
	err := c.BindJSON(&hardDiskDriveDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := hardDiskDriveApi.IHardDiskDriveService.Update(hardDiskDriveDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Hard disk drive updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (hardDiskDriveApi *HardDiskDriveAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := hardDiskDriveApi.IHardDiskDriveService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Hard disk drive deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}