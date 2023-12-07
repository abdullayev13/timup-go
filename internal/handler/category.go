package handler

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Category struct {
	Service *service.Service
}

func (h *Category) Create(c *gin.Context) {
	data := new(dtos.WorkCategory)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	//	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.Category.Create(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Category) Get(c *gin.Context) {
	parentId := 0
	if parentIdStr, ok := c.GetQuery("parent_id"); ok {
		parentId, _ = strconv.Atoi(parentIdStr)
	}

	res, err := h.Service.Category.GetByParentId(parentId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Category) Update(c *gin.Context) {
	data := new(dtos.WorkCategory)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	res, err := h.Service.Category.Update(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Category) Delete(c *gin.Context) {
	//	userId := c.GetInt(config.UserIdKeyFromAuthMw)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	err = h.Service.Category.DeleteById(id)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "ok")
}
