package dtos

import (
	"abdullayev13/timeup/internal/models"
	"abdullayev13/timeup/internal/utill"
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
	ParentId int    `json:"parent_id,omitempty"`
	Name     string `json:"name"`
}

type BusinessProfile struct {
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	CategoryId     int    `json:"category_id"`
	CategoryName   string `json:"category_name"`
	OfficeAddress  string `json:"office_address"`
	OfficeName     string `json:"office_name"`
	Experience     int    `json:"experience"`
	Bio            string `json:"bio"`
	DayOffs        string `json:"day_offs"`
	FollowersCount int    `json:"followers_count"`
	PostsCount     int    `json:"posts_count"`
}

type Booking struct {
	ID                int    `json:"id"`
	BusinessId        int    `json:"business_id"`
	ClientId          int    `json:"client_id"`
	BookingCategoryId *int   `json:"booking_category_id"`
	Date              string `json:"date"`
	Time              string `json:"time"`
	EndDate           string `json:"end_date"`
	EndTime           string `json:"end_time"`
}

type Following struct {
	ID         int       `json:"id"`
	BusinessId int       `json:"business_id"`
	FollowerId int       `json:"follower_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type BookingCategory struct {
	Id          int    `json:"id" form:"id"`
	BusinessId  int    `json:"business_id" form:"business_id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Duration    int    `json:"duration" form:"duration"`
	Price       int    `json:"price" form:"price"`
}

func (d *User) MapFromUser(m *models.User) *User {
	d.ID = m.ID
	d.FistName = m.FistName
	d.LastName = m.LastName
	d.UserName = m.UserName
	d.Address = m.Address
	d.PhoneNumber = m.PhoneNumber
	d.PhotoUrl = utill.PutMediaDomain(m.PhotoUrl)

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

func (d *BusinessProfile) MapFromModel(m *models.BusinessProfile) *BusinessProfile {
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

func (d *BusinessProfile) SetCategoryName(name string) {
	d.CategoryName = name
}

func (d *BusinessProfile) MapToModel() *models.BusinessProfile {
	m := new(models.BusinessProfile)

	m.ID = d.ID
	m.UserID = d.UserID
	m.WorkCategoryId = d.CategoryId
	m.OfficeAddress = d.OfficeAddress
	m.OfficeName = d.OfficeName
	m.Experience = d.Experience
	m.Bio = d.Bio
	m.DayOffs = d.DayOffs

	return m
}

func (d *WorkCategory) MapFromModel(m *models.WorkCategory) *WorkCategory {
	d.ID = m.ID
	d.Name = m.Name
	d.ParentId = m.ParentId

	return d
}

func (d *WorkCategory) MapToModel() *models.WorkCategory {
	m := new(models.WorkCategory)

	m.ID = d.ID
	m.Name = d.Name
	m.ParentId = d.ParentId

	return m
}

func (d *Booking) MapFromModel(m *models.Booking) *Booking {
	d.ID = m.ID
	d.BusinessId = m.BusinessId
	d.ClientId = m.ClientId
	d.BookingCategoryId = m.BookingCategoryId
	d.Date, d.Time = utill.Format(m.Date)
	d.EndDate, d.EndTime = utill.Format(m.EndTime)

	return d
}

func (d *Booking) MapToModel() *models.Booking {
	m := new(models.Booking)

	m.ID = d.ID
	m.BusinessId = d.BusinessId
	m.ClientId = d.ClientId
	m.BookingCategoryId = d.BookingCategoryId
	m.Date, _ = utill.Parse(d.Date, d.Time)
	m.EndTime, _ = utill.Parse(d.EndDate, d.EndTime)

	return m
}

func (d *Following) MapFromModel(m *models.Following) *Following {
	d.ID = m.ID
	d.BusinessId = m.BusinessId
	d.FollowerId = m.FollowerId
	d.CreatedAt = m.CreatedAt

	return d
}

func (d *Following) MapToModel() *models.Following {
	m := new(models.Following)

	m.ID = d.ID
	m.BusinessId = d.BusinessId
	m.FollowerId = d.FollowerId
	m.CreatedAt = d.CreatedAt

	return m
}

func (d *BookingCategory) MapToModel() *models.BookingCategory {
	m := new(models.BookingCategory)

	m.Id = d.Id
	m.BusinessId = d.BusinessId
	m.Name = d.Name
	m.Description = d.Description
	m.Duration = time.Duration(d.Duration) * time.Minute
	m.Price = d.Price

	return m
}

func (d *BookingCategory) MapFromModel(m *models.BookingCategory) *BookingCategory {
	d.Id = m.Id
	d.BusinessId = m.BusinessId
	d.Name = m.Name
	d.Description = m.Description
	d.Duration = int(m.Duration / time.Minute)
	d.Price = m.Price

	return d
}
