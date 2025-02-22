package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Quiz struct {
	ID        *uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Question  string     `gorm:"size:255" json:"question"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy string     `gorm:"size:50" json:"created_by"`
	UpdatedBy *string    `gorm:"size:50" json:"updated_by"`

	Choices []*QuizChoice `gorm:"foreignKey:QuizID;constraint:OnDelete:CASCADE" json:"choices"`
}

type QuizChoice struct {
	ID        *uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	QuizID    *uuid.UUID `gorm:"type:uuid;foreignKey" json:"quiz_id"`
	Choice    string     `gorm:"size:255" json:"choice"`
	Answer    string     `gorm:"size:255" json:"answer"`
	IsCorrect bool       `json:"is_correct"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy string     `gorm:"size:50" json:"created_by"`
	UpdatedBy *string    `gorm:"size:50" json:"updated_by"`
}

func (q *Quiz) GenUUID() {
	id, _ := uuid.NewV4()
	q.ID = &id
}

func (qc *QuizChoice) GenUUID() {
	id, _ := uuid.NewV4()
	qc.ID = &id
}

func (Quiz) TableName() string {
	return "quiz"
}

func (QuizChoice) TableName() string {
	return "quiz_choice"
}
