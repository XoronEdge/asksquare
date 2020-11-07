package domain

import (
	"context"
	"database/sql"
	"time"
)

// QaReport ...
type QaReport struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Reason       string `gorm:"" json:"reason" validate:"min=3" `
	Description  string `gorm:"" json:"description" validate:"min=6" `
	QaQuestionID uint
	UserID       uint
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    sql.NullTime `gorm:"index" json:"deleted_at"`
}

// QuestionHide ...
type QuestionHide struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	HideByUser uint `gorm:"-" json:"hideByUser" validate:"min=3" `
	QuestionID uint
	UserID     uint
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	DeletedAt  sql.NullTime `gorm:"index" json:"deleted_at"`
}

// QuestionAnswerLater ...
type QuestionAnswerLater struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	QuestionID uint
	UserID     uint
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	DeletedAt  sql.NullTime `gorm:"index" json:"deleted_at"`
}

// QuestionAnswerRequest ...
type QuestionAnswerRequest struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	RequestedByID uint         `json:"requestedById"`
	RequestedToID uint         `json:"requestedToId"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     sql.NullTime `gorm:"index" json:"deleted_at"`
}

// QaReportUsecase represent the User's usecases
type QaReportUsecase interface {
	Fetch(ctx context.Context) ([]*QaReport, error)
	GetByID(ctx context.Context, id uint) (*QaReport, error)
	Update(ctx context.Context, qar *QaReport) error
	Store(ctx context.Context, qar *QaReport) error
	Delete(ctx context.Context, id uint) error
}

// QaReportRepo represent the User's repository contract
type QaReportRepo interface {
	Fetch(ctx context.Context) ([]*QaReport, error)
	GetByID(ctx context.Context, id uint) (*QaReport, error)
	Update(ctx context.Context, qar *QaReport) error
	Store(ctx context.Context, qar *QaReport) error
	Delete(ctx context.Context, id uint) error
}
