package repo

import (
	"gorm.io/gorm"
)

type Repo struct {
	Users      *Users
	SmsCode    *SmsCode
	Business   *Business
	Category   *Category
	Booking    *Booking
	Following  *Following
	Post       *Post
	PostViewed *PostViewed
}

func New(DB *gorm.DB) *Repo {
	return &Repo{
		&Users{DB: DB},
		&SmsCode{DB},
		&Business{DB},
		&Category{DB},
		&Booking{DB},
		&Following{DB},
		&Post{DB},
		&PostViewed{DB},
	}
}
