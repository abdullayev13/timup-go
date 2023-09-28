package config

import "os"

var (
	Domain = ""
	Port   = ""
)

func LoadVarsFromEnv() {
	Domain = os.Getenv("DOMAIN")
	Port = os.Getenv("PORT")
}
