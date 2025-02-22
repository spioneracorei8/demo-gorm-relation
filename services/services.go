package services

import (
	"golang_relation/models"

	"github.com/gofrs/uuid"
)

type Exam interface {
	FetchListQuizChoices() ([]*models.Quiz, error)
	FetchQuizChoicesById(id *uuid.UUID) (*models.Quiz, error)
	CreateQuizChoices(*models.Quiz) (*models.Quiz, error)
}
