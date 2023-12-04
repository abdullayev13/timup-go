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
	DbHost              = "localhost"
	DbUsername          = "postgres"
	DbName              = "postgres"
	DbPassword          = "password"
	DbPort              = 5432
	EskizToken          = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDE2MDA4NTIsImlhdCI6MTY5OTAwODg1Miwicm9sZSI6InVzZXIiLCJzdWIiOiI0MTk3In0.bacrSAeF7eOGaXAa59oypdSMvLUnI3I_uTGmA6INpcw"
	ProfilePhotoDir     = "profilephoto"
	PostDir             = "post"
	S3BucketName        = "public-timeup-s3-media"
	FfmpegRunLimit      = 30
)

const (
	UserIdKeyFromAuthMw = "UserIdKeyFromAuthMw"
)

func LoadVarsFromEnv() {
	setIfExists(&Domain, "DOMAIN")
	setIfExists(&Port, "PORT")
	setIfExists(&JwtSignKey, "JWT_SIGN_KEY")

	setIfExists(&DbHost, "DB_HOST")
	setIfExists(&DbUsername, "DB_USERNAME")
	setIfExists(&DbName, "DB_DATABASE")
	setIfExists(&DbPassword, "DB_PASSWORD")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err == nil {
		DbPort = dbPort
	}
}

func setIfExists(ptr *string, key string) bool {
	value, ok := os.LookupEnv(key)
	if ok {
		*ptr = value
	}
	return ok
}

func IsDbBlank() bool {
	return DbHost == "" &&
		DbUsername == "" &&
		DbName == "" &&
		DbPassword == ""
}
