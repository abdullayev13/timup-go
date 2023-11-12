package handler

import (
	"abdullayev13/timeup/internal/config"
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Service *service.Service
}

func (h *Post) Create(c *gin.Context) {
	data := new(dtos.PostFile)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	data.Id = 0

	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.Post.Create(data, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Post) GetDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, errors.New("error when converting id: "+err.Error()))
		return
	}

	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.Post.GetDetail(id, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Post) GetList(c *gin.Context) {
	data := new(dtos.PostFilter)
	err := c.BindQuery(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data.BusinessId, err = strconv.Atoi(c.Param("business_id"))
	if err != nil {
		response.FailErr(c, errors.New("error when converting business_id: "+err.Error()))
		return
	}

	res, err := h.Service.Post.GetList(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Post) Update(c *gin.Context) {
	data := new(dtos.PostFile)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, errors.New("error when converting id: "+err.Error()))
		return
	}

	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.Post.Update(data, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Post) DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	err = h.Service.Post.DeleteById(id, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "ok")
}

// others

func (h *Post) GetListFollowed(c *gin.Context) {
	data := new(dtos.PostFilter)
	err := c.BindQuery(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.Post.GetListFollowed(data, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}
