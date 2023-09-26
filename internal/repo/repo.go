package repo

import (
	"gorm.io/gorm"
)

type Repo struct {
	Users   *Users
	SmsCode *SmsCode
}

func New(DB *gorm.DB) *Repo {
	return &Repo{Users: &Users{DB: DB}, SmsCode: &SmsCode{DB}}
}
