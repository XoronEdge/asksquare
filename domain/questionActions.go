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

// QaHide ...
type QaHide struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	HideByUser   uint `gorm:"" json:"hide_by_user" `
	QaQuestionID uint
	UserID       uint
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    sql.NullTime `gorm:"index" json:"deleted_at"`
}

// QaAnswerLater ...
type QaAnswerLater struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	QaQuestionID uint
	UserID       uint
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    sql.NullTime `gorm:"index" json:"deleted_at"`
}

// QaAnswerRequest ...
type QaAnswerRequest struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	RequestedByID uint         `json:"requestedById"`
	RequestedToID uint         `json:"requestedToId"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     sql.NullTime `gorm:"index" json:"deleted_at"`
}

// QaReportUsecase ...
type QaReportUsecase interface {
	Fetch(ctx context.Context) ([]*QaReport, error)
	GetByID(ctx context.Context, id uint) (*QaReport, error)
	Update(ctx context.Context, qar *QaReport) error
	Store(ctx context.Context, qar *QaReport) error
	Delete(ctx context.Context, id uint) error
}

// QaReportRepo ...
type QaReportRepo interface {
	Fetch(ctx context.Context) ([]*QaReport, error)
	GetByID(ctx context.Context, id uint) (*QaReport, error)
	Update(ctx context.Context, qar *QaReport) error
	Store(ctx context.Context, qar *QaReport) error
	Delete(ctx context.Context, id uint) error
}

// QaHideUsecase ...
type QaHideUsecase interface {
	Fetch(ctx context.Context) ([]*QaHide, error)
	GetByID(ctx context.Context, id uint) (*QaHide, error)
	Update(ctx context.Context, qah *QaHide) error
	Store(ctx context.Context, qah *QaHide) error
	Delete(ctx context.Context, id uint) error
}

//QaHideRepo ...
type QaHideRepo interface {
	Fetch(ctx context.Context) ([]*QaHide, error)
	GetByID(ctx context.Context, id uint) (*QaHide, error)
	Update(ctx context.Context, qah *QaHide) error
	Store(ctx context.Context, qah *QaHide) error
	Delete(ctx context.Context, id uint) error
}

// QaAnswerLaterUsecase ...
type QaAnswerLaterUsecase interface {
	Fetch(ctx context.Context) ([]*QaAnswerLater, error)
	GetByID(ctx context.Context, id uint) (*QaAnswerLater, error)
	Update(ctx context.Context, qal *QaAnswerLater) error
	Store(ctx context.Context, qal *QaAnswerLater) error
	Delete(ctx context.Context, id uint) error
}

//QaAnswerLaterRepo ...
type QaAnswerLaterRepo interface {
	Fetch(ctx context.Context) ([]*QaAnswerLater, error)
	GetByID(ctx context.Context, id uint) (*QaAnswerLater, error)
	Update(ctx context.Context, qal *QaAnswerLater) error
	Store(ctx context.Context, qal *QaAnswerLater) error
	Delete(ctx context.Context, id uint) error
}
