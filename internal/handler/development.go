package handler

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"github.com/gin-gonic/gin"
)

func (h *SmsCode) LastSentSms(c *gin.Context) {
	data := new(dtos.SendSmsReq)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	model, err := h.Service.SmsCode.LastSentSms(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, model)
}
