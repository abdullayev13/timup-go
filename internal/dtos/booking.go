package dtos

import (
	"abdullayev13/timeup/internal/models"
	"abdullayev13/timeup/internal/utill"
	"time"
)

type BookingFilter struct {
	Limit      int    `json:"limit" form:"limit"`
	Offset     int    `json:"offset" form:"offset"`
	Coming     bool   `json:"coming" form:"coming"`
	BusinessId int    `json:"business_id" form:"business_id"`
	ClientId   int    `json:"client_id" form:"client_id"`
	Date       string `json:"date" form:"date"`
}

type BookingMini struct {
	ID         int       `json:"id"`
	BusinessId int       `json:"business_id,omitempty"`
	ClientId   int       `json:"client_id"`
	Date       time.Time `json:"-"`
	DateJson   string    `json:"date"`
	Time       string    `json:"time"`

	End     time.Time `json:"-"`
	EndDate string    `json:"end_date"`
	EndTime string    `json:"end_time"`

	FistName    string `json:"fist_name"`
	LastName    string `json:"last_name"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	PhotoUrl    string `json:"photo_url"`
}

func (d *BookingMini) MapFromModel(m *models.Booking) *BookingMini {
	d.ID = m.ID
	d.BusinessId = m.BusinessId
	d.ClientId = m.ClientId
	d.DateJson, d.Time = utill.Format(m.Date)
	d.EndDate, d.EndTime = utill.Format(m.EndTime)

	return d
}

func (d *BookingMini) SetUser(user *models.User) *BookingMini {
	d.FistName = user.FistName
	d.LastName = user.LastName
	d.UserName = user.UserName
	d.PhoneNumber = user.PhoneNumber
	d.PhotoUrl = utill.PutMediaDomain(user.PhotoUrl)

	return d
}
