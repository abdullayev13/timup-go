package dtos

type BusinessMini struct {
	BusinessID int    `json:"business_id"`
	UserID     int    `json:"user_id"`
	Experience int    `json:"experience"`
	FistName   string `json:"fist_name"`
	LastName   string `json:"last_name"`
	PhotoUrl   string `json:"photo_url"`
}

type BusinessFilter struct {
	Limit      int
	Offset     int
	CategoryId int
}
