package domain

import (
	"context"
	"time"
)

// QaReply struct
type QaReply struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Description string `validate:"required" json:"description"`
	//QaCommentID        int16      `validate:"required" json:"qa_comment_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	//QaComment *QaComment `gorm:"foreignkey:QaCommentID" json:"qa_comment"`
}

// QaReplyUsecase represent the QaReply's usecases
type QaReplyUsecase interface {
	Fetch(ctx context.Context) ([]*QaReply, error)
	GetByID(ctx context.Context, id int64) (*QaReply, error)
	Update(ctx context.Context, update *QaReply) error
	Store(context.Context, *QaReply) error
	Delete(ctx context.Context, id int64) error
}

// QaReplyRepo represent the QaReply's repository
type QaReplyRepo interface {
	Fetch(ctx context.Context) ([]*QaReply, error)
	GetByID(ctx context.Context, id int64) (*QaReply, error)
	Update(ctx context.Context, update *QaReply) error
	Store(ctx context.Context, store *QaReply) error
	Delete(ctx context.Context, id int64) error
}
