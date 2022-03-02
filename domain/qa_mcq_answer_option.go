package domain

import (
	"context"
	"time"
)

// QaMcqAnswerOption struct
type QaMcqAnswerOption struct {
	ID uint `gorm:"primary_key" json:"id"`
	//QaMcqOptionID int16      `validate:"required" json:"qa_mcq_option_id"`
	//QaAnswerID int16      `validate:"required" json:"qa_answer_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	//QaMcqOption *QaMcqOption `gorm:"foreignkey:QaMcqOptionID" json:"qa_mcq_option"`
	//QaAnswer    *QaAnswer    `gorm:"foreignkey:QaAnswerID" json:"qa_answer"`
}

// QaMcqAnswerOptionUsecase represent the QaMcqAnswerOption's usecases
type QaMcqAnswerOptionUsecase interface {
	Fetch(ctx context.Context) ([]*QaMcqAnswerOption, error)
	GetByID(ctx context.Context, id int64) (*QaMcqAnswerOption, error)
	Update(ctx context.Context, update *QaMcqAnswerOption) error
	Store(context.Context, *QaMcqAnswerOption) error
	Delete(ctx context.Context, id int64) error
}

// QaMcqAnswerOptionRepo represent the QaMcqAnswerOption's repository
type QaMcqAnswerOptionRepo interface {
	Fetch(ctx context.Context) ([]*QaMcqAnswerOption, error)
	GetByID(ctx context.Context, id int64) (*QaMcqAnswerOption, error)
	Update(ctx context.Context, update *QaMcqAnswerOption) error
	Store(ctx context.Context, store *QaMcqAnswerOption) error
	Delete(ctx context.Context, id int64) error
}
