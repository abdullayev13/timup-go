package dtos

import "time"

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
