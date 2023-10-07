package handler

import (
	"abdullayev13/timeup/internal/config"
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"abdullayev13/timeup/internal/utill"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Service  *service.Service
	JwtToken *utill.TokenJWT
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

func (h *Auth) LogOut(c *gin.Context) {
	userId := c.GetInt(config.UserIdKeyFromAuthMw)
	println("log out : ", userId)
}
