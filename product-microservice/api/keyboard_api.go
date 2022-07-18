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


type KeyboardAPI struct {
	IKeyboardService service.IKeyboardService
}

func NewKeyboardAPI(keyboardService service.IKeyboardService) KeyboardAPI {
	return KeyboardAPI{IKeyboardService: keyboardService}
}

func (keyboardApi *KeyboardAPI) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }
	
	keyboards := keyboardApi.IKeyboardService.GetAll(page, pageSize)
	c.JSON(http.StatusOK, mapper.ToKeyboardDTOs(keyboards))
}

func (keyboardApi *KeyboardAPI) GetNumberOfRecords(c *gin.Context) {
	numberOfRecords := keyboardApi.IKeyboardService.GetNumberOfRecords()
	c.JSON(http.StatusOK, numberOfRecords)
}

func (keyboardApi *KeyboardAPI) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	keyboard, err := keyboardApi.IKeyboardService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToKeyboardDTO(keyboard))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (keyboardApi *KeyboardAPI) SearchByName(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	keyboards, err := keyboardApi.IKeyboardService.SearchByName(page, pageSize, c.Query("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToKeyboardDTOs(keyboards))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (keyboardApi *KeyboardAPI) GetNumberOfRecordsSearch(c *gin.Context) {
	numberOfRecords := keyboardApi.IKeyboardService.GetNumberOfRecordsSearch(c.Query("name"))
	c.JSON(http.StatusOK, numberOfRecords)
}

func (keyboardApi *KeyboardAPI) Filter(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	var filter filter.KeyboardFilter
	err = c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	keyboards, err := keyboardApi.IKeyboardService.Filter(page, pageSize, filter)
	
	if err == nil {
		c.JSON(http.StatusOK, mapper.ToKeyboardDTOs(keyboards))
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (keyboardApi *KeyboardAPI) GetNumberOfRecordsFilter(c *gin.Context) {
	var filter filter.KeyboardFilter
	err := c.BindJSON(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	numberOfRecords := keyboardApi.IKeyboardService.GetNumberOfRecordsFilter(filter)
	c.JSON(http.StatusOK, numberOfRecords)
}

func (keyboardApi *KeyboardAPI) GetManufacturers(c *gin.Context) {
	manufacturers := keyboardApi.IKeyboardService.GetManufacturers()
	c.JSON(http.StatusOK, manufacturers)
}

func (keyboardApi *KeyboardAPI) GetKeyboardConnectors(c *gin.Context) {
	keyboardConnectors := keyboardApi.IKeyboardService.GetKeyboardConnectors()
	c.JSON(http.StatusOK, keyboardConnectors)
}

func (keyboardApi *KeyboardAPI) GetKeyTypes(c *gin.Context) {
	keyTypes := keyboardApi.IKeyboardService.GetKeyTypes()
	c.JSON(http.StatusOK, keyTypes)
}

func (keyboardApi *KeyboardAPI) Create(c *gin.Context) {
	var keyboardDTO dto.KeyboardDTO
	err := c.BindJSON(&keyboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	keyboard := mapper.ToKeyboard(keyboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := keyboardApi.IKeyboardService.Create(keyboard)

	if msg == "" {
		c.JSON(http.StatusOK, "Keyboard added successfully.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (keyboardApi *KeyboardAPI) Update(c *gin.Context) {
	var keyboardDTO dto.KeyboardDTO
	err := c.BindJSON(&keyboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	msg := keyboardApi.IKeyboardService.Update(keyboardDTO)

	if msg == "" {
		c.JSON(http.StatusOK, "Keyboard updated successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, msg)
	} 
}

func (keyboardApi *KeyboardAPI) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	error := keyboardApi.IKeyboardService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, "Keyboard deleted successfully.")
	} else  {
		c.JSON(http.StatusBadRequest, error.Error())
	}
}