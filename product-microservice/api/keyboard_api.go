package api

import (
	"product/dto"
	"product/mapper"
	"product/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type KeyboardAPI struct {
	IKeyboardService service.IKeyboardService
}

func NewKeyboardAPI(keyboardService service.IKeyboardService) KeyboardAPI {
	return KeyboardAPI{IKeyboardService: keyboardService}
}

func (keyboardApi *KeyboardAPI) GetAll(c *gin.Context) {
	keyboards := keyboardApi.IKeyboardService.GetAll()
	c.JSON(http.StatusOK, gin.H{"keyboards": mapper.ToKeyboardDTOs(keyboards)})
}

func (keyboardApi *KeyboardAPI) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	keyboard, err := keyboardApi.IKeyboardService.GetById(id)
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"keyboard": mapper.ToKeyboardDTO(keyboard)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (keyboardApi *KeyboardAPI) GetByName(c *gin.Context) {
	keyboard, err := keyboardApi.IKeyboardService.GetByName(c.Param("name"))
	
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"keyboard": mapper.ToKeyboardDTO(keyboard)})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (keyboardApi *KeyboardAPI) Create(c *gin.Context) {
	var keyboardDTO dto.KeyboardDTO
	err := c.BindJSON(&keyboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	keyboard := mapper.ToKeyboard(keyboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := keyboardApi.IKeyboardService.Create(keyboard)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Keyboard stored successfully."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}

func (keyboardApi *KeyboardAPI) Update(c *gin.Context) {
	var keyboardDTO dto.KeyboardDTO
	err := c.BindJSON(&keyboardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	error := keyboardApi.IKeyboardService.Update(keyboardDTO)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Keyboard updated successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	} 
}

func (keyboardApi *KeyboardAPI) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	error := keyboardApi.IKeyboardService.Delete(id)

	if error == nil {
		c.JSON(http.StatusOK, gin.H{"msg": "Keyboard deleted successfully."})
	} else  {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}
}