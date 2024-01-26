package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Variables struct {
	DbSource            string
	PasetoSymetricToken string
	PasetoTokenDuration time.Duration
}

func NewVars() *Variables {

	godotenv.Load(".env")

	// duration, _ := time.ParseDuration(os.Getenv("PASETO_TOKEN_DURATION"))

	return &Variables{
		DbSource:            os.Getenv("DB_URL"),
		PasetoSymetricToken: os.Getenv(("PASETO_SYMETRIC_TOKEN")),
		PasetoTokenDuration: time.Minute,
	}

}
