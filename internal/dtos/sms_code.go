package dtos

import "abdullayev13/timeup/internal/models"

type SendSmsReq struct {
	PhoneNumber string `json:"phone_number"`
}

type VerifySmsReq struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}
type VerifySmsRes struct {
	Register bool
	Token    string
	User     *models.User
}
