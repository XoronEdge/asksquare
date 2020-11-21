package domain

import (
	"context"
	"time"
)

// QaCmOption struct
type QaCmOption struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Title string `validate:"required" json:"title"`
	//QaQuestionID       int16      `validate:"required" json:"qa_question_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	//QaQuestion *QaQuestion `gorm:"foreignkey:QaQuestionID" json:"qa_question"`
	//CmAnswer QaCmAnswer `validate:"-" gorm:"foreignkey:QaCmOptionID" json:"cm_answer"`
}

// QaCmOptionUsecase represent the QaCmOption's usecases
type QaCmOptionUsecase interface {
	Fetch(ctx context.Context) ([]*QaCmOption, error)
	GetByID(ctx context.Context, id int64) (*QaCmOption, error)
	Update(ctx context.Context, update *QaCmOption) error
	Store(context.Context, *QaCmOption) error
	Delete(ctx context.Context, id int64) error
}

// QaCmOptionRepo represent the QaCmOption's repository
type QaCmOptionRepo interface {
	Fetch(ctx context.Context) ([]*QaCmOption, error)
	GetByID(ctx context.Context, id int64) (*QaCmOption, error)
	Update(ctx context.Context, update *QaCmOption) error
	Store(ctx context.Context, store *QaCmOption) error
	Delete(ctx context.Context, id int64) error
}
