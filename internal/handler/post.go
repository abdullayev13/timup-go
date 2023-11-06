package handler

import (
	"abdullayev13/timeup/internal/config"
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
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

	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.Post.Create(data, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Post) GetDetail(c *gin.Context) {
	//data := new(dtos.PostFilter)
	//err := c.BindQuery(data)
	//if err != nil {
	//	response.FailErr(c, err)
	//	return
	//}
	//
	//data.BusinessId, err = strconv.Atoi(c.Param("business_id"))
	//
	//res, err := h.Service.Post.GetList(data)
	//if err != nil {
	//	response.FailErr(c, err)
	//	return
	//}
	//
	//response.Success(c, res)
}

func (h *Post) GetList(c *gin.Context) {
	data := new(dtos.PostFilter)
	err := c.BindQuery(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data.BusinessId, err = strconv.Atoi(c.Param("business_id"))

	res, err := h.Service.Post.GetList(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Post) Update(c *gin.Context) {
	//data := new(dtos.Post)
	//err := c.Bind(data)
	//if err != nil {
	//	response.FailErr(c, err)
	//	return
	//}
	//
	//data.ClientId = c.GetInt(config.UserIdKeyFromAuthMw)
	//
	//res, err := h.Service.Post.Create(data, 0)
	//if err != nil {
	//	response.FailErr(c, err)
	//	return
	//}
	//
	//response.Success(c, res)
}

func (h *Post) DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	err = h.Service.Post.DeleteById(id)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "ok")
}
