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


type SolidStateDriveAPI struct {
	ISolidStateDriveService service.ISolidStateDriveService
}

func NewSolidStateDriveAPI(solidStateDriveService service.ISolidStateDriveService) SolidStateDriveAPI {
	return SolidStateDriveAPI{ISolidStateDriveService: solidStateDriveService}
}

func (solidStateDriveApi *SolidStateDriveAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	solidStateDrives := solidStateDriveApi.ISolidStateDriveService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToSolidStateDriveDTOs(solidStateDrives))
}

func (solidStateDriveApi *SolidStateDriveAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	solidStateDrive, err := solidStateDriveApi.ISolidStateDriveService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToSolidStateDriveDTO(solidStateDrive))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (solidStateDriveApi *SolidStateDriveAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	solidStateDrives, err := solidStateDriveApi.ISolidStateDriveService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToSolidStateDriveDTOs(solidStateDrives))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (solidStateDriveApi *SolidStateDriveAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.SolidStateDriveFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	solidStateDrives, err := solidStateDriveApi.ISolidStateDriveService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToSolidStateDriveDTOs(solidStateDrives))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (solidStateDriveApi *SolidStateDriveAPI) GetCapacities(c *gin.Context) {
	capacities := solidStateDriveApi.ISolidStateDriveService.GetCapacities()
	c.JSON(http.StatusOK, capacities)
}

func (solidStateDriveApi *SolidStateDriveAPI) GetForms(c *gin.Context) {
	forms := solidStateDriveApi.ISolidStateDriveService.GetForms()
	c.JSON(http.StatusOK, forms)
}

func (solidStateDriveApi *SolidStateDriveAPI) GetManufacturers(c *gin.Context) {
	manufacturers := solidStateDriveApi.ISolidStateDriveService.GetManufacturers()
	c.JSON(http.StatusOK, manufacturers)
}

func (solidStateDriveApi *SolidStateDriveAPI) GetMaxSequentialReads(c *gin.Context) {
	maxSequentialReads := solidStateDriveApi.ISolidStateDriveService.GetMaxSequentialReads()
	c.JSON(http.StatusOK, maxSequentialReads)
}

func (solidStateDriveApi *SolidStateDriveAPI) GetMaxSequentialWrites(c *gin.Context) {
	maxSequentialWrites := solidStateDriveApi.ISolidStateDriveService.GetMaxSequentialWrites()
	c.JSON(http.StatusOK, maxSequentialWrites)
}

func (solidStateDriveApi *SolidStateDriveAPI) Create(c *gin.Context) {
	var solidStateDriveDTO dto.SolidStateDriveDTO
	err := c.BindJSON(&solidStateDriveDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	solidStateDrive := mapper.ToSolidStateDrive(solidStateDriveDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := solidStateDriveApi.ISolidStateDriveService.Create(solidStateDrive)

	if error == nil {
		c.JSON(http.StatusOK, "Solid state drive stored successfully.")
	} else {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}

func (solidStateDriveApi *SolidStateDriveAPI) Update(c *gin.Context) {
	var solidStateDriveDTO dto.SolidStateDriveDTO
	err := c.BindJSON(&solidStateDriveDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := solidStateDriveApi.ISolidStateDriveService.Update(solidStateDriveDTO)

	if error == nil {
		c.JSON(http.StatusOK, "Solid state drive updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	} 
}

func (solidStateDriveApi *SolidStateDriveAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := solidStateDriveApi.ISolidStateDriveService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "Solid state drive deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}