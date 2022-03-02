package domain

import (
	"context"
	"time"
)

// QaCmPrompt struct
type QaCmPrompt struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Title string `validate:"required" json:"title"`
	//QaQuestionID       int16      `validate:"required" json:"qa_question_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	//QaQuestion *QaQuestion `gorm:"foreignkey:QaQuestionID" json:"qa_question"`
	//CmAnswer QaCmAnswer `validate:"-" gorm:"foreignkey:QaCmPromptID" json:"cm_answer"`
}

// QaCmPromptUsecase represent the QaCmPrompt's usecases
type QaCmPromptUsecase interface {
	Fetch(ctx context.Context) ([]*QaCmPrompt, error)
	GetByID(ctx context.Context, id int64) (*QaCmPrompt, error)
	Update(ctx context.Context, update *QaCmPrompt) error
	Store(context.Context, *QaCmPrompt) error
	Delete(ctx context.Context, id int64) error
}

// QaCmPromptRepo represent the QaCmPrompt's repository
type QaCmPromptRepo interface {
	Fetch(ctx context.Context) ([]*QaCmPrompt, error)
	GetByID(ctx context.Context, id int64) (*QaCmPrompt, error)
	Update(ctx context.Context, update *QaCmPrompt) error
	Store(ctx context.Context, store *QaCmPrompt) error
	Delete(ctx context.Context, id int64) error
}
