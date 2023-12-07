package handler

import (
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"abdullayev13/timeup/internal/utill"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Dev struct {
	Service *service.Service
}

func (h *Dev) EskizSetData(c *gin.Context) {
	type EskizData struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}
	data := new(EskizData)
	err := c.ShouldBind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	if data.Email == "" {
		c.ShouldBindQuery(data)
		if data.Email == "" {
			marshal, _ := json.Marshal(data)
			response.Fail(c, "data not found: "+string(marshal))
			return
		}
	}

	err = utill.SetEskizData(data.Email, data.Password)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "done")
}

func (h *Dev) EskizRefreshToken(c *gin.Context) {
	err := utill.EskizRefreshToken()
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "done")
}
