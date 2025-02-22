package main

import (
	"golang_relation/helper"
	my_logger "golang_relation/logger"
	"golang_relation/models"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("Error loading .env file: %s", err.Error())
	}
	PSQL_CONNECTION := helper.GetENV("PSQL_CONNECTION", "")
	gormLogger := my_logger.GormLogger()
	db, err := gorm.Open(postgres.Open(PSQL_CONNECTION), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logrus.Fatalf("Error connecting to database: %s", err.Error())
	}
	db.Migrator().AutoMigrate(
		&models.Quiz{},
		&models.QuizChoice{},
	)
}
