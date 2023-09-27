package dtos

import (
	"abdullayev13/timeup/internal/models"
	"mime/multipart"
)

type RegisterReq struct {
	ProfilePhoto *multipart.FileHeader `form:"profile_photo" json:"profile_photo"`
	PhotoUrl     string                `form:"-" json:"-"`

	FistName    string `form:"fist_name" json:"fist_name"`
	LastName    string `form:"last_name" json:"last_name"`
	Password    string `form:"password" json:"password"`
	UserName    string `form:"user_name" json:"user_name"`
	PhoneNumber string `form:"phone_number" json:"phone_number"`
	Address     string `form:"address" json:"address"`
}
type RegisterRes struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}
