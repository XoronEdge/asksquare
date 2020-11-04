package domain

import (
	"database/sql"
	"time"
)

// Question ...
type Question struct {
	ID                    uint   `gorm:"primaryKey" json:"id"`
	Title                 string `gorm:"-" json:"title" validate:"min=7" `
	Description           string `gorm:"-" json:"description" validate:"min=7" `
	Question              string `json:"question" validate:"required,min=7" `
	QuestionReports       []QuestionReports
	QuestionHides         []QuestionHides
	QuestionAnswerLater   []QuestionAnswerLater
	QuestionAnswerRequest []QuestionAnswerRequest
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             time.Time    `json:"updated_at"`
	DeletedAt             sql.NullTime `gorm:"index" json:"deleted_at"`
}
