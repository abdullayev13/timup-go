package handler

import (
	"abdullayev13/timeup/internal/service"
)

type Handlers struct {
	SmsCode *SmsCode
	Auth    *Auth
}

func New(serv *service.Service) *Handlers {
	return &Handlers{SmsCode: &SmsCode{serv}, Auth: &Auth{serv}}
}
