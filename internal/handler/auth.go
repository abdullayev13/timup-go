package handler

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Service *service.Service
}

func (h *Auth) Register(c *gin.Context) {
	data := new(dtos.RegisterReq)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	res, err := h.Service.Users.Register(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}
