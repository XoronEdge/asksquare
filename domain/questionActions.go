package domain

import (
	"database/sql"
	"time"
)

// QuestionReports ...
type QuestionReports struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Reason      string `gorm:"-" json:"reason" validate:"min=3" `
	Description string `gorm:"-" json:"description" validate:"min=6" `
	QuestionID  uint
	UserID      uint
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `gorm:"index" json:"deleted_at"`
}

// QuestionHides ...
type QuestionHides struct {
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
