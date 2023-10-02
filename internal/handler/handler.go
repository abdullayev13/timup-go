package handler

import (
	"abdullayev13/timeup/internal/service"
	"abdullayev13/timeup/internal/utill"
)

type Handlers struct {
	SmsCode  *SmsCode
	Auth     *Auth
	User     *User
	Business *Business
	Category *Category
}

func New(serv *service.Service, jwtToken *utill.TokenJWT) *Handlers {
	return &Handlers{
		&SmsCode{serv},
		&Auth{serv, jwtToken},
		&User{serv, jwtToken},
		&Business{serv},
		&Category{serv},
	}
}
