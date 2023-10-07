package handler

import (
	"abdullayev13/timeup/internal/config"
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"abdullayev13/timeup/internal/utill"
	"github.com/gin-gonic/gin"
)

type User struct {
	Service  *service.Service
	JwtToken *utill.TokenJWT
}

func (h *User) UserMe(c *gin.Context) {
	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.Users.GetUserBusiness(userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *User) EditMe(c *gin.Context) {
	userId := c.GetInt(config.UserIdKeyFromAuthMw)
	data := new(dtos.User)

	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	data.ID = userId

	res, err := h.Service.Users.Update(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *User) EditPhoto(c *gin.Context) {
	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	data := new(dtos.PhotoReq)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	data.UserId = userId

	res, err := h.Service.Users.EditPhoto(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)

}

func (h *User) DeleteMe(c *gin.Context) {
	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	err := h.Service.Users.DeleteById(userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "ok")
}
