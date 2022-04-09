package api

import (
	"product/dto"
	"product/mapper"
	"product/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type SolidStateDriveAPI struct {
	ISolidStateDriveService service.ISolidStateDriveService
}

func NewSolidStateDriveAPI(solidStateDriveService service.ISolidStateDriveService) SolidStateDriveAPI {
	return SolidStateDriveAPI{ISolidStateDriveService: solidStateDriveService}
}

func (solidStateDriveApi *SolidStateDriveAPI) GetAll(c *gin.Context) {
	solidStateDrives := solidStateDriveApi.ISolidStateDriveService.GetAll()
	c.JSON(http.StatusOK, gin.H{"solid_state_drives": mapper.ToSolidStateDriveDTOs(solidStateDrives)})
}

func (solidStateDriveApi *SolidStateDriveAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	solidStateDrive, err := solidStateDriveApi.ISolidStateDriveService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"solid_state_drive": mapper.ToSolidStateDriveDTO(solidStateDrive)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (solidStateDriveApi *SolidStateDriveAPI) GetByName(c *gin.Context) {
	solidStateDrive, err := solidStateDriveApi.ISolidStateDriveService.GetByName(c.Param("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"solid_state_drive": mapper.ToSolidStateDriveDTO(solidStateDrive)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (solidStateDriveApi *SolidStateDriveAPI) Create(c *gin.Context) {
	var solidStateDriveDTO dto.SolidStateDriveDTO
	err := c.BindJSON(&solidStateDriveDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	solidStateDrive := mapper.ToSolidStateDrive(solidStateDriveDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := solidStateDriveApi.ISolidStateDriveService.Create(solidStateDrive)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Solid state drive stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (solidStateDriveApi *SolidStateDriveAPI) Update(c *gin.Context) {
	var solidStateDriveDTO dto.SolidStateDriveDTO
	err := c.BindJSON(&solidStateDriveDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := solidStateDriveApi.ISolidStateDriveService.Update(solidStateDriveDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Solid state drive updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (solidStateDriveApi *SolidStateDriveAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := solidStateDriveApi.ISolidStateDriveService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Solid state drive deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}