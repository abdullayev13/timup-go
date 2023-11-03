package utill

import (
	"abdullayev13/timeup/internal/config"
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func SendSmsCode(phoneNumber string, code int) {
	SendSms(phoneNumber, fmt.Sprintf(
		"TimeUp \n verification code: %d", code))
}

func SendSms(phoneNumber, smsMsg string) {

	url := "notify.eskiz.uz/api/message/sms/send"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("mobile_phone", phoneNumber)
	_ = writer.WriteField("message", smsMsg)
	_ = writer.WriteField("from", "4546")
	_ = writer.WriteField("callback_url", "http://0000.uz/test.php")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	req.Header.Add("Authorization", "Bearer "+config.EskizToken)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
