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

		photoUrl = config.Domain + photoUrl
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
	res.User = new(dtos.User)
	res.User.MapFromUser(user)

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

	userDto := new(dtos.User)
	userDto.MapFromUser(userModel)

	return userDto, nil
}

func (s *Users) Update(dto *dtos.User) (*dtos.User, error) {
	model := dto.MapToUser()
	model, err := s.Repo.Users.Update(model)
	if err != nil {
		return nil, err
	}

	dto.MapFromUser(model)

	return dto, nil
}

func (s *Users) DeleteById(id int) error {
	return s.Repo.Users.DeleteById(id)
}

/*




















 */
