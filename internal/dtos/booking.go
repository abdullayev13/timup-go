package dtos

type BookingFilter struct {
	Limit      int
	Offset     int
	Coming     bool
	BusinessId int
	ClientId   int
	Date       string
}
