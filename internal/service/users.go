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
	Repo *repo.Repo
}

func (s *Users) SignUp(sign dtos.Sign) (*models.User, error) {
	exists := s.Repo.Users.ExistsByUsername(sign.UserName)
	if exists {
		return nil, errors.New("username exists")
	}
	user := models.User{UserName: sign.UserName, Password: sign.Password, FistName: sign.Name}
	//s.DB.Create(&user)
	return &user, nil
}

func (s *Users) LogIn(sign dtos.Sign) (string, error) {
	exists := s.Repo.Users.ExistsByUsername(sign.UserName)
	if !exists {
		return "", errors.New("username or password wrong")
	}
	//var userId int
	//s.DB.Model(&moduls.User{}).
	//	Select("id").
	//	Where("user_name = ?", sign.UserName).
	//	Find(&userId)
	return "token", nil
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
	{
		res.User.ID = user.ID
		res.User.FistName = user.FistName
		res.User.LastName = user.LastName
		res.User.UserName = user.UserName
		res.User.Address = user.Address
		res.User.PhoneNumber = user.PhoneNumber
		res.User.PhotoUrl = user.PhotoUrl
	}
	//TODO make token
	res.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	return res, nil
}
