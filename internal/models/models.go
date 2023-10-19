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
	FistName    string
	LastName    string
	Password    string
	UserName    string `gorm:"index:idx_username,unique;not null"`
	PhoneNumber string `gorm:"index:idx_phone_number,unique;not null"`
	Address     string
	PhotoUrl    string
	BirthDate   *time.Time
}

type WorkCategory struct {
	ID       int
	ParentId int    `gorm:"index:idx_parent_id__name,unique;"`
	Name     string `gorm:"index:idx_parent_id__name,unique;not null"`
}

type BusinessProfile struct {
	ID             int
	UserID         int
	WorkCategoryId int
	OfficeAddress  string
	OfficeName     string
	Experience     int
	Bio            string
	DayOffs        string
}

type Booking struct {
	ID         int
	BusinessId int
	ClientId   int
	Date       time.Time
}

type Following struct {
	ID         int
	BusinessId int `gorm:"index:idx_following,unique;not null"`
	FollowerId int `gorm:"index:idx_following,unique;not null"`
	CreatedAt  time.Time
}
