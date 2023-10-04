package dtos

type SendSmsReq struct {
	PhoneNumber string `json:"phone_number"`
}

type VerifySmsReq struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}
type VerifySmsRes struct {
	Register bool          `json:"register"`
	Token    string        `json:"token"`
	User     *UserBusiness `json:"user"`
}
