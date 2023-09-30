package handler

import (
	"abdullayev13/timeup/internal/config"
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"github.com/gin-gonic/gin"
)

type Business struct {
	Service *service.Service
}

func (h *Business) Create(c *gin.Context) {
	data := new(dtos.BusinessProfile)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data.UserID = c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.Business.Create(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Business) GetMe(c *gin.Context) {
	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.Business.GetByUserId(userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}
