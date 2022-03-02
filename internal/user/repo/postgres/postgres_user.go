package postgres

import (
	"context"

	"github.com/XoronEdge/asksquare/domain"
	"github.com/jinzhu/gorm"
)

type userRepo struct {
	db *gorm.DB
}

// NewUserRepo  ...
func NewUserRepo(db *gorm.DB) domain.UserRepo {
	return &userRepo{db}
}

//Fetch ...
func (u *userRepo) Fetch(ctx context.Context) ([]*domain.User, error) {
	users := []*domain.User{}
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

//Store ...
func (u *userRepo) Store(ctx context.Context, newUser *domain.User) (err error) {
	err = u.db.Create(&newUser).Error
	if err != nil {
		return
	}
	return
}

//GetById ...
func (u *userRepo) GetByID(ctx context.Context, id int64) (user *domain.User, err error) {
	user = &domain.User{}
	err = u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return
}

//Update ...
func (u *userRepo) Update(ctx context.Context, user *domain.User) (err error) {
	err = u.db.Save(&user).Error
	if err != nil {
		return err
	}
	return
}

//Delete ...
func (u *userRepo) Delete(ctx context.Context, id int64) (err error) {
	err = u.db.Delete(&domain.User{}, id).Error
	return
}
