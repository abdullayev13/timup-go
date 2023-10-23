package dtos

import (
	"abdullayev13/timeup/internal/models"
	"abdullayev13/timeup/internal/utill"
)

type BusinessMini struct {
	BusinessID int    `json:"business_id"`
	UserID     int    `json:"user_id"`
	Experience int    `json:"experience"`
	FistName   string `json:"fist_name"`
	LastName   string `json:"last_name"`
	PhotoUrl   string `json:"photo_url"`
	Followed   bool   `json:"followed"`
}

type BusinessFilter struct {
	Limit      int `json:"limit" form:"limit"`
	Offset     int `json:"offset" form:"offset"`
	CategoryId int `json:"category_id" form:"category_id"`
	UserId     int `json:"user_id" form:"user_id"`
}

type BusinessData struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	CategoryId    int    `json:"category_id,omitempty"`
	CategoryName  string `json:"category_name"`
	OfficeAddress string `json:"office_address"`
	OfficeName    string `json:"office_name"`
	Experience    int    `json:"experience"`
	Bio           string `json:"bio"`
	DayOffs       string `json:"day_offs"`
	FistName      string `json:"fist_name"`
	LastName      string `json:"last_name"`
	UserName      string `json:"user_name"`
	PhoneNumber   string `json:"phone_number"`
	Address       string `json:"address"`
	PhotoUrl      string `json:"photo_url"`
}

type BusinessFullData struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	CategoryId    int    `json:"category_id"`
	CategoryName  string `json:"category_name"`
	OfficeAddress string `json:"office_address"`
	OfficeName    string `json:"office_name"`
	Experience    int    `json:"experience"`
	Bio           string `json:"bio"`
	DayOffs       string `json:"day_offs"`
	FistName      string `json:"fist_name"`
	LastName      string `json:"last_name"`
	UserName      string `json:"user_name"`
	PhoneNumber   string `json:"phone_number"`
	Address       string `json:"address"`
	PhotoUrl      string `json:"photo_url"`
}

func (d *BusinessData) MapFromModel(m *models.BusinessProfile) *BusinessData {
	d.ID = m.ID
	d.UserID = m.UserID
	d.CategoryId = m.WorkCategoryId
	d.OfficeAddress = m.OfficeAddress
	d.OfficeName = m.OfficeName
	d.Experience = m.Experience
	d.Bio = m.Bio
	d.DayOffs = m.DayOffs

	return d
}

func (d *BusinessData) SetCategoryName(name string) {
	d.CategoryName = name
}

func (d *BusinessData) SetUser(user *models.User) {
	d.FistName = user.FistName
	d.LastName = user.LastName
	d.UserName = user.UserName
	d.PhoneNumber = user.PhoneNumber
	d.Address = user.Address
	d.PhotoUrl = utill.PutMediaDomain(user.PhotoUrl)
}
