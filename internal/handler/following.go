package handler

import (
	"abdullayev13/timeup/internal/config"
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Following struct {
	Service *service.Service
}

func (h *Following) Create(c *gin.Context) {
	businessId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	userId := c.GetInt(config.UserIdKeyFromAuthMw)
	data := new(dtos.Following)
	data.BusinessId = businessId
	data.FollowerId = userId

	res, err := h.Service.Following.Create(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Following) DeleteByFollower(c *gin.Context) {
	businessId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	err = h.Service.Following.DeleteByFollower(businessId, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "ok")
}

func (h *Following) GetBusinessList(c *gin.Context) {
	data := new(dtos.FollowedFilter)
	err := c.BindQuery(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	data.FollowerId = c.GetInt(config.UserIdKeyFromAuthMw)

	list, err := h.Service.Following.GetBusinessList(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, list)
}
