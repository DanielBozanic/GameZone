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


type PowerSupplyUnitAPI struct {
	IPowerSupplyUnitService service.IPowerSupplyUnitService
}

func NewPowerSupplyUnitAPI(powerSupplyUnitService service.IPowerSupplyUnitService) PowerSupplyUnitAPI {
	return PowerSupplyUnitAPI{IPowerSupplyUnitService: powerSupplyUnitService}
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	powerSupplyUnits := powerSupplyUnitApi.IPowerSupplyUnitService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToPowerSupplyUnitDTOs(powerSupplyUnits))
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	powerSupplyUnit, err := powerSupplyUnitApi.IPowerSupplyUnitService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToPowerSupplyUnitDTO(powerSupplyUnit))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	powerSupplyUnits, err := powerSupplyUnitApi.IPowerSupplyUnitService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToPowerSupplyUnitDTOs(powerSupplyUnits))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.PowerSupplyUnitFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	powerSupplyUnits, err := powerSupplyUnitApi.IPowerSupplyUnitService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToPowerSupplyUnitDTOs(powerSupplyUnits))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) GetManufacturers(c *gin.Context) {
	manufacturers := powerSupplyUnitApi.IPowerSupplyUnitService.GetManufacturers()
	c.JSON(http.StatusOK, manufacturers)
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) GetPowers(c *gin.Context) {
	powers := powerSupplyUnitApi.IPowerSupplyUnitService.GetPowers()
	c.JSON(http.StatusOK, powers)
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) GetTypes(c *gin.Context) {
	types := powerSupplyUnitApi.IPowerSupplyUnitService.GetTypes()
	c.JSON(http.StatusOK, types)
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) GetFormFactors(c *gin.Context) {
	formFactors := powerSupplyUnitApi.IPowerSupplyUnitService.GetFormFactors()
	c.JSON(http.StatusOK, formFactors)
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) Create(c *gin.Context) {
	var powerSupplyUnitDTO dto.PowerSupplyUnitDTO
	err := c.BindJSON(&powerSupplyUnitDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	powerSupplyUnit := mapper.ToPowerSupplyUnit(powerSupplyUnitDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := powerSupplyUnitApi.IPowerSupplyUnitService.Create(powerSupplyUnit)

	if error == nil {
		c.JSON(http.StatusOK, "PSU stored successfully.")
	} else {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) Update(c *gin.Context) {
	var powerSupplyUnitDTO dto.PowerSupplyUnitDTO
	err := c.BindJSON(&powerSupplyUnitDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := powerSupplyUnitApi.IPowerSupplyUnitService.Update(powerSupplyUnitDTO)

	if error == nil {
		c.JSON(http.StatusOK, "PSU updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	} 
}

func (powerSupplyUnitApi *PowerSupplyUnitAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := powerSupplyUnitApi.IPowerSupplyUnitService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "PSU deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}