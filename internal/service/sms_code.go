package service

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/models"
	"abdullayev13/timeup/internal/repo"
	"abdullayev13/timeup/internal/utill"
	"errors"
	"strconv"
	"time"
)

type SmsCode struct {
	Repo *repo.Repo
}

func (s *SmsCode) SendSms(data *dtos.SendSmsReq) error {
	code := utill.Random6DigNum()
	model := new(models.SmsCode)
	model.PhoneNumber = data.PhoneNumber
	model.Code = strconv.Itoa(code)
	model.SentAt = time.Now()

	var err error
	model, err = s.Repo.SmsCode.Create(model)
	if err != nil {
		return err
	}

	return nil
}

func (s *SmsCode) LastSentSms(data *dtos.SendSmsReq) (*models.SmsCode, error) {
	model, err := s.Repo.SmsCode.LastByPhoneNumber(data.PhoneNumber)

	return model, err
}

func (s *SmsCode) VerifySmsCode(data *dtos.VerifySmsReq) (*dtos.VerifySmsRes, error) {
	model, err := s.Repo.SmsCode.LastByPhoneNumber(data.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if model.Code != data.Code {
		return nil, errors.New("do not match")
	}
	model.Verified = true
	_, err = s.Repo.SmsCode.Update(model)
	if err != nil {
		return nil, err
	}

	res := new(dtos.VerifySmsRes)

	exists := s.Repo.Users.ExistsByPhoneNumber(data.PhoneNumber)
	if !exists {
		res.Register = true
		return res, nil
	}

	user, err := s.Repo.Users.GetByPhoneNumber(data.PhoneNumber)
	if err != nil {
		return nil, err
	}

	res.User = user
	//TODO make token
	res.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	return res, nil
}
