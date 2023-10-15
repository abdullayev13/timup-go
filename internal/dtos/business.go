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
	Limit      int `json:"limit" form:"limit"`
	Offset     int `json:"offset" form:"offset"`
	CategoryId int `json:"category_id" form:"category_id"`
}
