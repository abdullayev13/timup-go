package utill

import (
	"abdullayev13/timeup/internal/config"
	"fmt"
	"github.com/realtemirov/eskizuz"
)

func SendSmsCode(phoneNumber string, code int) {
	SendSms(phoneNumber, fmt.Sprintf(
		"TimeUp \n verification code: %d", code))
}

func SendSms(phoneNumber, smsMsg string) {
	if len(phoneNumber) > 12 {
		phoneNumber = phoneNumber[1:]
	}

	sms := &eskizuz.SMS{
		MobilePhone: phoneNumber,
		Message:     smsMsg,
		From:        "4546",
		CallbackURL: "https://eskiz.uz",
	}

	// Sending message
	result, err := config.Eskiz.Send(sms)
	_, _ = result, err
}
