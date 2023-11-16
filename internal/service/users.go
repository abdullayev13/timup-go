package service

import (
	"abdullayev13/timeup/internal/config"
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
	if !utill.ValidPhoneNumber(data.PhoneNumber) {
		return nil, errors.New("PhoneNumber is not valid")
	}

	smscode, err := s.Repo.SmsCode.LastByPhoneNumber(data.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if !smscode.Verified {
		return nil, errors.New("phone number is not verified")
	}

	var photoUrl string
	if data.ProfilePhoto != nil {
		photoUrl, err = utill.Upload(data.ProfilePhoto, config.ProfilePhotoDir)
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

		bd, err := utill.ParseDate(data.BirthDate)
		if err != nil {
			return nil, errors.New("err parsing BirthDate: " + err.Error())
		}
		user.BirthDate = &bd
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

	userModel.PhotoUrl = utill.PutMediaDomain(userModel.PhotoUrl)

	userDto := new(dtos.User)
	userDto.MapFromUser(userModel)

	return userDto, nil
}

func (s *Users) GetUserBusiness(userId int) (*dtos.UserBusiness, error) {
	fullData, err := s.Repo.Business.GetProfileByUserId(userId)

	if err != nil || fullData.UserID == 0 {
		model, err := s.Repo.Users.GetById(userId)
		if err != nil {
			return nil, err
		}
		dto := new(dtos.UserBusiness).MapFromModel(model)
		return dto, nil
	}

	fullData.PhotoUrl = utill.PutMediaDomain(fullData.PhotoUrl)

	dto := new(dtos.UserBusiness)
	{
		dto.ID = fullData.UserID
		dto.FistName = fullData.FistName
		dto.LastName = fullData.LastName
		dto.UserName = fullData.UserName
		dto.PhoneNumber = fullData.PhoneNumber
		dto.Address = fullData.Address
		dto.PhotoUrl = fullData.PhotoUrl
		dto.FollowingCount = fullData.FollowingCount
	}
	if fullData.ID == 0 {
		return dto, nil
	}

	dto.Business = new(dtos.BusinessProfile)
	{
		dto.Business.ID = fullData.ID
		dto.Business.UserID = fullData.UserID
		dto.Business.CategoryId = fullData.CategoryId
		dto.Business.CategoryName = fullData.CategoryName
		dto.Business.OfficeAddress = fullData.OfficeAddress
		dto.Business.OfficeName = fullData.OfficeName
		dto.Business.Experience = fullData.Experience
		dto.Business.Bio = fullData.Bio
		dto.Business.DayOffs = fullData.DayOffs
		dto.Business.FollowersCount = fullData.FollowersCount
	}

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

func (s *Users) EditPhoto(data *dtos.PhotoReq) (*dtos.PhotoRes, error) {
	user, err := s.Repo.Users.GetById(data.UserId)
	if err != nil {
		return nil, err
	}

	var photoUrl = data.PhotoUrl
	if data.ProfilePhoto != nil {
		photoUrl, err = utill.Upload(data.ProfilePhoto, "profilephoto")
		if err != nil {
			return nil, err
		}
	}

	user.PhotoUrl = photoUrl
	user, err = s.Repo.Users.Update(user)
	if err != nil {
		return nil, err
	}

	res := new(dtos.PhotoRes)
	res.PhotoUrl = utill.PutMediaDomain(photoUrl)

	return res, nil
}
