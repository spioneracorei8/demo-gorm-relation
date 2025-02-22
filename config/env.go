package config

import (
	"golang_relation/helper"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	PSQL_CONNECTION string
	APP_PORT        string
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("Error loading .env file: %s", err.Error())
	}
	PSQL_CONNECTION = helper.GetENV("PSQL_CONNECTION", "")
	APP_PORT = helper.GetENV("APP_PORT", "")
}
