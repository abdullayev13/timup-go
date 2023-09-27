package models

import (
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
	ID          int
	FistName    string `json:"fist_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"-"`
	UserName    string `gorm:"index:idx_username,unique;not null" json:"user_name"`
	PhoneNumber string `gorm:"index:idx_phone_number,unique;not null" json:"phone_number"`
	Address     string `json:"address"`
	PhotoUrl    string `json:"photo_url"`
}
type WorkCategory struct {
	ID       int    `json:"id"`
	ParentId int    `gorm:"index:idx_parent_id__name,unique;" json:"parent_id"`
	Name     string `gorm:"index:idx_parent_id__name,unique;not null" json:"name"`
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
