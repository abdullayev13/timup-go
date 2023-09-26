package handler

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"github.com/gin-gonic/gin"
)

type SmsCode struct {
	Service *service.Service
}

func (h *SmsCode) SendSms(c *gin.Context) {
	data := new(dtos.SendSmsReq)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	err = h.Service.SmsCode.SendSms(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, nil)
}

func (h *SmsCode) VerifySmsCode(c *gin.Context) {
	data := new(dtos.VerifySmsReq)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	res, err := h.Service.SmsCode.VerifySmsCode(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}
	response.Success(c, res)
}
