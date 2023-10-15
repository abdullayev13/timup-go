package handler

import (
	"abdullayev13/timeup/internal/config"
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
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

func (h *Business) GetByCategory(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data := new(dtos.BusinessFilter)
	err = c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data.CategoryId = categoryId
	res, err := h.Service.Business.GetByGetByCategory(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Business) UpdateMe(c *gin.Context) {
	data := new(dtos.BusinessProfile)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data.UserID = c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.Business.Update(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Business) DeleteMe(c *gin.Context) {
	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	err := h.Service.Business.DeleteByUserId(userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "ok")
}
