package server

import (
	"fmt"
	exam_controller "golang_relation/controllers/exam"
	my_logger "golang_relation/logger"
	"golang_relation/routes"
	exam_service "golang_relation/services/exam"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	PSQL_CONNECTION string
	APP_PORT        string
}

func connectDatabase(PSQL_CONNECTION string) *gorm.DB {
	gormLogger := my_logger.GormLogger()
	db, err := gorm.Open(postgres.Open(PSQL_CONNECTION), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logrus.Fatalf("Error connecting to database: %s", err.Error())
	}
	return db
}

func (s Server) Start() {
	app := fiber.New(
		fiber.Config{
			AppName: "Golang Relation",
		},
	)
	db := connectDatabase(s.PSQL_CONNECTION)

	examSvcs := exam_service.NewPsqlExamServiceImpl(db)
	examCtrl := exam_controller.NewExamCtrlImpl(examSvcs)

	router := routes.NewRoute(app)

	router.NewExamRoute(examCtrl)

	if err := app.Listen(fmt.Sprintf("%s%s", ":", s.APP_PORT)); err != nil {
		logrus.Errorf("Error listening app port %s : %s", s.APP_PORT, err.Error())
	}
}
