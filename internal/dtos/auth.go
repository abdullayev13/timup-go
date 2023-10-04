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
	User  *UserBusiness `json:"user"`
	Token string        `json:"token"`
}

type UserBusiness struct {
	ID          int              `json:"id"`
	FistName    string           `json:"fist_name"`
	LastName    string           `json:"last_name"`
	UserName    string           `json:"user_name"`
	PhoneNumber string           `json:"phone_number"`
	Address     string           `json:"address"`
	PhotoUrl    string           `json:"photo_url"`
	Business    *BusinessProfile `json:"business"`
}

func (d *UserBusiness) MapFromModel(m *models.User) *UserBusiness {
	d.ID = m.ID
	d.FistName = m.FistName
	d.LastName = m.LastName
	d.UserName = m.UserName
	d.Address = m.Address
	d.PhoneNumber = m.PhoneNumber
	d.PhotoUrl = m.PhotoUrl

	return d
}

func (d *UserBusiness) MapToModel() *models.User {
	m := new(models.User)

	m.ID = d.ID
	m.FistName = d.FistName
	m.LastName = d.LastName
	m.UserName = d.UserName
	m.Address = d.Address
	m.PhoneNumber = d.PhoneNumber
	m.PhotoUrl = d.PhotoUrl

	return m
}
