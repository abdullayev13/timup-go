package dtos

type FollowedFilter struct {
	FollowerId int
	Limit      int
	Offset     int
}

type BusinessLI struct {
	ID            int    `json:"id"`
	CategoryName  string `json:"category_name"`
	OfficeAddress string `json:"office_address"`
	OfficeName    string `json:"office_name"`
	Bio           string `json:"bio"`
	DayOffs       string `json:"day_offs"`
	FistName      string `json:"fist_name"`
	LastName      string `json:"last_name"`
	UserName      string `json:"user_name"`
	PhoneNumber   string `json:"phone_number"`
	PhotoUrl      string `json:"photo_url"`
}
