package handler

import (
	"abdullayev13/timeup/internal/config"
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BookingCategory struct {
	Service *service.Service
}

func (h *BookingCategory) Create(c *gin.Context) {
	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	data := new(dtos.BookingCategory)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	//	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.BookingCategory.Create(data, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *BookingCategory) GetList(c *gin.Context) {
	businessId, err := strconv.Atoi(c.Param("business_id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	res, err := h.Service.BookingCategory.GetByBusinessId(businessId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *BookingCategory) Delete(c *gin.Context) {
	userId := c.GetInt(config.UserIdKeyFromAuthMw)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	err = h.Service.BookingCategory.DeleteById(id, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "ok")
}
