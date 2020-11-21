package domain

import (
	"context"
	"time"
)

// QaMcqOption struct
type QaMcqOption struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Title string `validate:"required" json:"title"`
	//QaQuestionID       int16      `validate:"required" json:"qa_question_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	//QaQuestion *QaQuestion `gorm:"foreignkey:QaQuestionID" json:"qa_question"`
	//McqAnswerOption QaMcqAnswerOption `validate:"-" gorm:"foreignkey:QaMcqOptionID" json:"mcq_answer_option"`
}

// QaMcqOptionUsecase represent the QaMcqOption's usecases
type QaMcqOptionUsecase interface {
	Fetch(ctx context.Context) ([]*QaMcqOption, error)
	GetByID(ctx context.Context, id int64) (*QaMcqOption, error)
	Update(ctx context.Context, update *QaMcqOption) error
	Store(context.Context, *QaMcqOption) error
	Delete(ctx context.Context, id int64) error
}

// QaMcqOptionRepo represent the QaMcqOption's repository
type QaMcqOptionRepo interface {
	Fetch(ctx context.Context) ([]*QaMcqOption, error)
	GetByID(ctx context.Context, id int64) (*QaMcqOption, error)
	Update(ctx context.Context, update *QaMcqOption) error
	Store(ctx context.Context, store *QaMcqOption) error
	Delete(ctx context.Context, id int64) error
}
