package service

import (
	"abdullayev13/timeup/internal/repo"
)

type Service struct {
	Users   *Users
	SmsCode *SmsCode
}

func New(repository *repo.Repo) *Service {
	return &Service{Users: &Users{repository}, SmsCode: &SmsCode{repository}}
}
