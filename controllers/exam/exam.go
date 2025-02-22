package exam

import (
	"golang_relation/controllers"
	"golang_relation/models"
	"golang_relation/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

type ExamCtrl struct {
	examService services.Exam
}

func NewExamCtrlImpl(examService services.Exam) controllers.Exam {
	return &ExamCtrl{
		examService: examService,
	}
}

func (h *ExamCtrl) FetchListQuizChoics(f *fiber.Ctx) error {
	quizChoices, err := h.examService.FetchListQuizChoices()
	if err != nil {
		return f.Status(http.StatusInternalServerError).JSON(fiber.Map{})
	}
	return f.Status(http.StatusOK).JSON(fiber.Map{
		"quiz_choices": quizChoices,
	})
}

func (h *ExamCtrl) FetchQuizChoicesById(f *fiber.Ctx) error {
	quizId := uuid.FromStringOrNil(f.Params("quiz_id"))
	quizze, err := h.examService.FetchQuizChoicesById(&quizId)
	if err != nil {
		return f.Status(http.StatusInternalServerError).JSON(fiber.Map{})
	}
	return f.Status(http.StatusOK).JSON(fiber.Map{
		"quiz_choices": quizze,
	})
}

func (h *ExamCtrl) CreateQuizChoices(f *fiber.Ctx) error {
	var quizChoices *models.Quiz
	if err := f.BodyParser(&quizChoices); err != nil {
		return f.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	quizChoices.GenUUID()
	for _, c := range quizChoices.Choices {
		c.GenUUID()
	}

	response, err := h.examService.CreateQuizChoices(quizChoices)
	if err != nil {
		return f.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return f.Status(http.StatusOK).JSON(fiber.Map{
		"quiz_choices": response,
	})
}
