package domain

import (
	"context"
	"database/sql"
	"time"
)

// User ...
type User struct {
	ID                    uint   `gorm:"primaryKey" json:"id"`
	Username              string `gorm:"size(5);not null;uniqueIndex" json:"username" validate:"required,min=5"`
	Email                 string `gorm:"not null;uniqueIndex" json:"email" validate:"required,email"`
	HashPassword          string `gorm:"size(8);not null" json:"hashPassword"`
	Password              string `gorm:"-" json:"password" validate:"required,min=8" `
	Gender                string `json:"gender" validate:"min=3" `
	FirstName             string `json:"first_name" validate:"min=3" `
	LastName              string `json:"last_name" validate:"min=3" `
	Phone                 string `json:"phone" validate:"required,min=7" `
	QuestionAnswerLater   []QuestionAnswerLater
	QaReport              []QaReport
	QuestionHides         []QaHide
	QuestionAnswerRequest []QuestionAnswerRequest
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             time.Time    `json:"updated_at"`
	DeletedAt             sql.NullTime `gorm:"index" json:"deleted_at"`
}

// UserUsecase represent the User's usecases
type UserUsecase interface {
	Fetch(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Update(ctx context.Context, ar *User) error
	Store(context.Context, *User) error
	Delete(ctx context.Context, id int64) error
}

// UserRepo represent the User's repository contract
type UserRepo interface {
	Fetch(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Update(ctx context.Context, ar *User) error
	Store(ctx context.Context, a *User) error
	Delete(ctx context.Context, id int64) error
}
