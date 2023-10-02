package service

import (
	"abdullayev13/timeup/internal/repo"
	"abdullayev13/timeup/internal/utill"
)

type Service struct {
	Users    *Users
	SmsCode  *SmsCode
	Business *Business
	Category *Category
}

func New(repository *repo.Repo, jwtToken *utill.TokenJWT) *Service {
	return &Service{
		&Users{repository, jwtToken},
		&SmsCode{repository, jwtToken},
		&Business{repository},
		&Category{repository},
	}
}
