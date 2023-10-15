package repo

import (
	"gorm.io/gorm"
)

type Repo struct {
	Users    *Users
	SmsCode  *SmsCode
	Business *Business
	Category *Category
	Booking  *Booking
}

func New(DB *gorm.DB) *Repo {
	return &Repo{
		&Users{DB: DB},
		&SmsCode{DB},
		&Business{DB},
		&Category{DB},
		&Booking{DB},
	}
}
