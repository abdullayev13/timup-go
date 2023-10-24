package config

import (
	"os"
	"strconv"
	"time"
)

var (
	Domain              = ""
	Port                = ""
	JwtSignKey          = "secret"
	JwtExpiringDuration = time.Hour * 24 * 100_000
	DbHost              = ""
	DbUsername          = ""
	DbName              = ""
	DbPassword          = ""
	DbPort              = 5432
)

const (
	UserIdKeyFromAuthMw = "UserIdKeyFromAuthMw"
)

func LoadVarsFromEnv() {
	Domain = os.Getenv("DOMAIN")
	Port = os.Getenv("PORT")
	JwtSignKey = os.Getenv("JWT_SIGN_KEY")

	DbHost = os.Getenv("DB_HOST")
	DbUsername = os.Getenv("DB_USERNAME")
	DbName = os.Getenv("DB_DATABASE")
	DbPassword = os.Getenv("DB_PASSWORD")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err == nil {
		DbPort = dbPort
	}
}
