package repo

import (
	"gorm.io/gorm"
)

type Repo struct {
	Users    *Users
	SmsCode  *SmsCode
	Business *Business
	Category *Category
}

func New(DB *gorm.DB) *Repo {
	return &Repo{
		&Users{DB: DB},
		&SmsCode{DB},
		&Business{DB},
		&Category{DB},
	}
}
