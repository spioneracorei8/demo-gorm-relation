package routes

import (
	"golang_relation/controllers"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	f *fiber.App
}

func NewRoute(f *fiber.App) *Route {
	return &Route{f: f}
}

func (r *Route) NewExamRoute(ctrl controllers.Exam) {
	r.f.Get("/v1/exam/quiz-choices", ctrl.FetchListQuizChoics)
	r.f.Get("/v1/exam/quiz-choices/:quiz_id", ctrl.FetchQuizChoicesById)
	r.f.Post("/v1/exam/quiz-choices", ctrl.CreateQuizChoices)
}
