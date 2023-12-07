package utill

import (
	"abdullayev13/timeup/internal/config"
	"github.com/realtemirov/eskizuz"
)

func SetEskizData(email, password string) error {

	eskiz, err := eskizuz.GetToken(&eskizuz.Auth{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return err
	}

	config.Eskiz = eskiz
	return nil
}

func EskizRefreshToken() error {
	return config.Eskiz.RefreshToken()
}
