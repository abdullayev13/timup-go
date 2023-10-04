package service

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/models"
	"abdullayev13/timeup/internal/repo"
	"abdullayev13/timeup/internal/utill"
	"errors"
)

type Users struct {
	Repo     *repo.Repo
	JwtToken *utill.TokenJWT
}

func (s *Users) Register(data *dtos.RegisterReq) (*dtos.RegisterRes, error) {
	// TODO data.Address checking
	smscode, err := s.Repo.SmsCode.LastByPhoneNumber(data.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if !smscode.Verified {
		return nil, errors.New("phone number is not verified")
	}

	var photoUrl string
	if data.ProfilePhoto != nil {
		photoUrl, err = utill.Upload(data.ProfilePhoto, "profilephoto")
		if err != nil {
			return nil, err
		}
	}

	user := new(models.User)
	{
		user.PhotoUrl = photoUrl

		user.FistName = data.FistName
		user.LastName = data.LastName
		user.Password = data.Password
		user.UserName = data.UserName
		user.PhoneNumber = data.PhoneNumber
		user.Address = data.Address
	}

	user, err = s.Repo.Users.Create(user)
	if err != nil {
		return nil, err
	}

	res := new(dtos.RegisterRes)
	res.User, err = s.GetUserBusiness(user.ID)
	if err != nil {
		return nil, err
	}

	res.Token, err = s.JwtToken.Generate(user.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Users) GetById(userId int) (*dtos.User, error) {
	userModel, err := s.Repo.Users.GetById(userId)
	if err != nil {
		return nil, err
	}

	userModel.PhotoUrl = utill.PutDomain(userModel.PhotoUrl)

	userDto := new(dtos.User)
	userDto.MapFromUser(userModel)

	return userDto, nil
}

func (s *Users) GetUserBusiness(userId int) (*dtos.UserBusiness, error) {
	userModel, err := s.Repo.Users.GetById(userId)
	if err != nil {
		return nil, err
	}

	userModel.PhotoUrl = utill.PutDomain(userModel.PhotoUrl)

	dto := new(dtos.UserBusiness)
	dto.MapFromModel(userModel)

	businessModel, err := s.Repo.Business.GetByUserId(userId)
	if err != nil {
		return dto, nil
	}

	dto.Business = new(dtos.BusinessProfile)
	dto.Business.MapFromModel(businessModel)

	return dto, nil
}

func (s *Users) Update(dto *dtos.User) (*dtos.User, error) {
	model := dto.MapToUser()
	orgModel, err := s.Repo.Users.GetById(model.ID)
	if err != nil {
		return nil, err
	}

	model.PhoneNumber = orgModel.PhoneNumber
	if model.PhotoUrl == "" {
		model.PhotoUrl = orgModel.PhotoUrl
	}

	model, err = s.Repo.Users.Update(model)
	if err != nil {
		return nil, err
	}

	dto.MapFromUser(model)

	return dto, nil
}

func (s *Users) DeleteById(id int) error {
	err := s.Repo.Business.DeleteByUserId(id)
	if err != nil {
		return err
	}

	return s.Repo.Users.DeleteById(id)
}
