package service

import (
	"abdullayev13/timeup/internal/repo"
	"abdullayev13/timeup/internal/utill"
)

type Service struct {
	Users     *Users
	SmsCode   *SmsCode
	Business  *Business
	Category  *Category
	Booking   *Booking
	Following *Following
	Post      *Post
}

func New(repository *repo.Repo, jwtToken *utill.TokenJWT) *Service {
	users := &Users{repository, jwtToken}
	return &Service{
		users,
		&SmsCode{repository, jwtToken, users},
		&Business{repository},
		&Category{repository},
		&Booking{repository},
		&Following{repository},
		&Post{repository},
	}
}
