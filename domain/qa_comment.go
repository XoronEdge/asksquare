package domain

import (
	"context"
	"time"
)

// QaComment struct
type QaComment struct {
	ID                uint   `gorm:"primary_key" json:"id"`
	Description       string `validate:"required" json:"description"`
	TotalRepliesCount int    `json:"total_replies_count"`
	//QaAnswerID       int16      `validate:"required" json:"qa_answer_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	//QaAnswer *QaAnswer `gorm:"foreignkey:QaAnswerID" json:"qa_answer"`
	// Replies []QaReply `validate:"-" gorm:"foreignkey:QaCommentID" json:"replies"`
}

// QaCommentUsecase represent the QaComment's usecases
type QaCommentUsecase interface {
	Fetch(ctx context.Context) ([]*QaComment, error)
	GetByID(ctx context.Context, id int64) (*QaComment, error)
	Update(ctx context.Context, update *QaComment) error
	Store(context.Context, *QaComment) error
	Delete(ctx context.Context, id int64) error
}

// QaCommentRepo represent the QaComment's repository
type QaCommentRepo interface {
	Fetch(ctx context.Context) ([]*QaComment, error)
	GetByID(ctx context.Context, id int64) (*QaComment, error)
	Update(ctx context.Context, update *QaComment) error
	Store(ctx context.Context, store *QaComment) error
	Delete(ctx context.Context, id int64) error
}
