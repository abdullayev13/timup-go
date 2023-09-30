package service

import (
	"abdullayev13/timeup/internal/repo"
	"abdullayev13/timeup/internal/utill"
)

type Service struct {
	Users   *Users
	SmsCode *SmsCode
}

func New(repository *repo.Repo, jwtToken *utill.TokenJWT) *Service {
	return &Service{
		Users:   &Users{repository, jwtToken},
		SmsCode: &SmsCode{repository, jwtToken},
	}
}
