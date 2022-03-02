package domain

import (
	"context"
	"time"
)

// QaAnswer struct
type QaAnswer struct {
	ID                 uint   `gorm:"primary_key" json:"id"`
	Description        string `validate:"required" json:"description"`
	Explaination       string `json:"explaination"`
	TotalCommentsCount int    `json:"total_comments_count"`
	//QaQuestionID       int16      `validate:"required" json:"qa_question_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	//QaQuestion *QaQuestion `gorm:"foreignkey:QaQuestionID" json:"qa_question"`
	// Comments         []QaComment         `validate:"-" gorm:"foreignkey:QaAnswerID" json:"comments"`
	// FibAnswers       []QaFibAnswer       `validate:"-" gorm:"foreignkey:QaAnswerID" json:"fib_answers"`
	// McqAnswerOptions []QaMcqAnswerOption `validate:"-" gorm:"foreignkey:QaAnswerID" json:"mcq_answer_options"`
	// CmAnswers        []QaCmAnswer        `validate:"-" gorm:"foreignkey:QaAnswerID" json:"cm_answers"`
}

// QaAnswerUsecase represent the QaAnswer's usecases
type QaAnswerUsecase interface {
	Fetch(ctx context.Context) ([]*QaAnswer, error)
	GetByID(ctx context.Context, id int64) (*QaAnswer, error)
	Update(ctx context.Context, update *QaAnswer) error
	Store(context.Context, *QaAnswer) error
	Delete(ctx context.Context, id int64) error
}

// QaAnswerRepo represent the QaAnswer's repository
type QaAnswerRepo interface {
	Fetch(ctx context.Context) ([]*QaAnswer, error)
	GetByID(ctx context.Context, id int64) (*QaAnswer, error)
	Update(ctx context.Context, update *QaAnswer) error
	Store(ctx context.Context, store *QaAnswer) error
	Delete(ctx context.Context, id int64) error
}
