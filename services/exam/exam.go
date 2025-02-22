package exam

import (
	"golang_relation/models"
	"golang_relation/services"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type psqlExamService struct {
	db *gorm.DB
}

func NewPsqlExamServiceImpl(db *gorm.DB) services.Exam {
	return &psqlExamService{
		db: db,
	}
}

func (r *psqlExamService) FetchListQuizChoices() ([]*models.Quiz, error) {
	var quizzes []*models.Quiz

	if err := r.db.
		Preload("Choices").
		Find(&quizzes).
		Error; err != nil {
		return nil, err
	}

	return quizzes, nil
}

func (r *psqlExamService) FetchQuizChoicesById(id *uuid.UUID) (*models.Quiz, error) {
	var quiz *models.Quiz

	if err := r.db.
		Preload("Choices").
		Where("id = ?", id).
		Find(&quiz).
		Error; err != nil {
		return nil, err
	}

	return quiz, nil
}

func (r *psqlExamService) CreateQuizChoices(quizChoices *models.Quiz) (*models.Quiz, error) {
	var quiz *models.Quiz

	if err := r.db.
		Create(&quizChoices).
		Preload("Choices").
		Where("id = ?", quizChoices.ID).
		Last(&quiz).
		Error; err != nil {
		return nil, err
	}

	return quiz, nil
}
