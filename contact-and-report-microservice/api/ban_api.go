package api

import (
	"contact-and-report/dto"
	"contact-and-report/mapper"
	"contact-and-report/service"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type BanAPI struct {
	IBanService service.IBanService
}

func NewBanAPI(banService service.IBanService) BanAPI {
	return BanAPI{IBanService: banService}
}

func (banApi *BanAPI) GetUserBanHistory(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	bans := banApi.IBanService.GetUserBanHistory(userId)
	c.JSON(http.StatusOK, mapper.ToBanDTOs(bans))
}

func (banApi *BanAPI) IsUserBanned(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
    if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
    }

	isUserBanned := banApi.IBanService.IsUserBanned(userId)
	c.JSON(http.StatusOK, isUserBanned)
}

func (banApi *BanAPI) AddBan(c *gin.Context) {
	var banDTO dto.BanDTO
	err := c.BindJSON(&banDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ban := mapper.ToBan(banDTO)
	msg := banApi.IBanService.AddBan(ban)
	if msg == "" {
		c.JSON(http.StatusOK, "User is banned.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}

func (banApi *BanAPI) SendEmailToBannedUser(c *gin.Context) {
	var banDTO dto.BanDTO
	err := c.BindJSON(&banDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ban := mapper.ToBan(banDTO)
	msg := banApi.IBanService.SendEmailToBannedUser(ban)
	if msg == "" {
		c.JSON(http.StatusOK, "Email about ban is sent.")
	} else {
		c.JSON(http.StatusBadRequest, msg)
	}
}