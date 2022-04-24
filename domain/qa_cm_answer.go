package domain

import (
	"context"
	"time"
)

// QaCmAnswer struct
type QaCmAnswer struct {
	ID uint `gorm:"primary_key" json:"id"`

	// QaCmPromptID int16      `validate:"required" json:"qa_cm_prompt_id"`
	// QaCmOptionID int16      `validate:"required" json:"qa_cm_option_id"`
	// QaAnswerID   int16      `validate:"required" json:"qa_answer_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	// QaCmPrompt *QaCmPrompt `gorm:"foreignkey:QaCmPromptID" json:"qa_cm_prompt"`
	// QaCmOption *QaCmOption `gorm:"foreignkey:QaCmOptionID" json:"qa_cm_option"`
	// QaAnswer   *QaAnswer   `gorm:"foreignkey:QaAnswerID" json:"qa_answer"`
}

// QaCmAnswerUsecase represent the QaCmAnswer's usecases
type QaCmAnswerUsecase interface {
	Fetch(ctx context.Context) ([]*QaCmAnswer, error)
	GetByID(ctx context.Context, id int64) (*QaCmAnswer, error)
	Update(ctx context.Context, update *QaCmAnswer) error
	Store(context.Context, *QaCmAnswer) error
	Delete(ctx context.Context, id int64) error
}

// QaCmAnswerRepo represent the QaCmAnswer's repository
type QaCmAnswerRepo interface {
	Fetch(ctx context.Context) ([]*QaCmAnswer, error)
	GetByID(ctx context.Context, id int64) (*QaCmAnswer, error)
	Update(ctx context.Context, update *QaCmAnswer) error
	Store(ctx context.Context, store *QaCmAnswer) error
	Delete(ctx context.Context, id int64) error
}
