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


type RamAPI struct {
	IRamService service.IRamService
}

func NewRamAPI(ramService service.IRamService) RamAPI {
	return RamAPI{IRamService: ramService}
}

func (ramApi *RamAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	rams := ramApi.IRamService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToRamDTOs(rams))
}

func (ramApi *RamAPI) GetNumberOfRecords(c *gin.Context) {
	numberOfRecords := ramApi.IRamService.GetNumberOfRecords()
	c.JSON(http.StatusOK, numberOfRecords)
}

func (ramApi *RamAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	ram, err := ramApi.IRamService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToRamDTO(ram))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (ramApi *RamAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	rams, err := ramApi.IRamService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToRamDTOs(rams))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (ramApi *RamAPI) GetNumberOfRecordsSearch(c *gin.Context) {
	numberOfRecords := ramApi.IRamService.GetNumberOfRecordsSearch(c.Query("name"))
	c.JSON(http.StatusOK, numberOfRecords)
}

func (ramApi *RamAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.RAMFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rams, err := ramApi.IRamService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToRamDTOs(rams))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (ramApi *RamAPI) GetNumberOfRecordsFilter(c *gin.Context) {
	var filter filter.RAMFilter
	err := c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	numberOfRecords := ramApi.IRamService.GetNumberOfRecordsFilter(filter)
	c.JSON(http.StatusOK, numberOfRecords)
}

func (ramApi *RamAPI) GetManufacturers(c *gin.Context) {
	manufacturers := ramApi.IRamService.GetManufacturers()
	c.JSON(http.StatusOK, manufacturers)
}

func (ramApi *RamAPI) GetCapacities(c *gin.Context) {
	capacities := ramApi.IRamService.GetCapacities()
	c.JSON(http.StatusOK, capacities)
}

func (ramApi *RamAPI) GetMemoryTypes(c *gin.Context) {
	memoryTypes := ramApi.IRamService.GetMemoryTypes()
	c.JSON(http.StatusOK, memoryTypes)
}


func (ramApi *RamAPI) GetSpeeds(c *gin.Context) {
	speeds := ramApi.IRamService.GetSpeeds()
	c.JSON(http.StatusOK, speeds)
}

func (ramApi *RamAPI) Create(c *gin.Context) {
	var ramDTO dto.RamDTO
	err := c.BindJSON(&ramDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ram := mapper.ToRam(ramDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := ramApi.IRamService.Create(ram)

	if error == nil {
		c.JSON(http.StatusOK, "RAM stored successfully.")
	} else {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}

func (ramApi *RamAPI) Update(c *gin.Context) {
	var ramDTO dto.RamDTO
	err := c.BindJSON(&ramDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	error := ramApi.IRamService.Update(ramDTO)

	if error == nil {
		c.JSON(http.StatusOK, "RAM updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	} 
}

func (ramApi *RamAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := ramApi.IRamService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "RAM deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}