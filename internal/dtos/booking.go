package dtos

type BookingFilter struct {
	Limit      int    `json:"limit" form:"limit"`
	Offset     int    `json:"offset" form:"offset"`
	Coming     bool   `json:"coming" form:"coming"`
	BusinessId int    `json:"business_id" form:"business_id"`
	ClientId   int    `json:"client_id" form:"client_id"`
	Date       string `json:"date" form:"date"`
}
