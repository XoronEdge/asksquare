package domain

import (
	"context"
	"time"
)

// QaFibAnswer struct
type QaFibAnswer struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Answer   string `validate:"required,max=50" json:"answer"`
	OptionNo int    `validate:"required" json:"option_no"`
	//QaAnswerID       int16      `validate:"required" json:"qa_answer_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	//QaAnswer *QaAnswer `gorm:"foreignkey:QaAnswerID" json:"qa_answer"`
}

// QaFibAnswerUsecase represent the QaFibAnswer's usecases
type QaFibAnswerUsecase interface {
	Fetch(ctx context.Context) ([]*QaFibAnswer, error)
	GetByID(ctx context.Context, id int64) (*QaFibAnswer, error)
	Update(ctx context.Context, update *QaFibAnswer) error
	Store(context.Context, *QaFibAnswer) error
	Delete(ctx context.Context, id int64) error
}

// QaFibAnswerRepo represent the QaFibAnswer's repository
type QaFibAnswerRepo interface {
	Fetch(ctx context.Context) ([]*QaFibAnswer, error)
	GetByID(ctx context.Context, id int64) (*QaFibAnswer, error)
	Update(ctx context.Context, update *QaFibAnswer) error
	Store(ctx context.Context, store *QaFibAnswer) error
	Delete(ctx context.Context, id int64) error
}
