package dtos

import (
	"abdullayev13/timeup/internal/models"
	"time"
)

type SmsCode struct {
	ID          int
	PhoneNumber string
	Code        string
	SentAt      time.Time
	Verified    bool
}

type User struct {
	ID          int    `json:"id"`
	FistName    string `json:"fist_name"`
	LastName    string `json:"last_name"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	PhotoUrl    string `json:"photo_url"`
}
type WorkCategory struct {
	ID       int    `json:"id"`
	ParentId int    `json:"parent_id"`
	Name     string `json:"name"`
}
type BusinessProfile struct {
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	WorkCategoryId int    `json:"work_category_id"`
	OfficeAddress  string `json:"office_address"`
	OfficeName     string `json:"office_name"`
	Experience     int    `json:"experience"`
	Bio            string `json:"bio"`
	DayOffs        string `json:"day_offs"`
}
type Booking struct {
	ID        int       `json:"id"`
	ServiceId int       `json:"service_id"`
	ClientId  int       `json:"client_id"`
	Date      time.Time `json:"date"`
}

func (d *User) MapFromUser(m *models.User) *User {
	d.ID = m.ID
	d.FistName = m.FistName
	d.LastName = m.LastName
	d.UserName = m.UserName
	d.Address = m.Address
	d.PhoneNumber = m.PhoneNumber
	d.PhotoUrl = m.PhotoUrl

	return d
}

func (d *User) MapToUser() *models.User {
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
