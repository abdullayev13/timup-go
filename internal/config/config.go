package config

import (
	"os"
	"time"
)

var (
	Domain              = ""
	Port                = ""
	JwtSignKey          = "secret"
	JwtExpiringDuration = time.Hour * 24 * 100_000
)

func LoadVarsFromEnv() {
	Domain = os.Getenv("DOMAIN")
	Port = os.Getenv("PORT")
	JwtSignKey = os.Getenv("JWT_SIGN_KEY")
}
