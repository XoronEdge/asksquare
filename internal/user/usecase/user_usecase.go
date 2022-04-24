package usecase

import (
	"context"
	"time"

	"github.com/XoronEdge/asksquare/domain"
	validator "github.com/go-playground/validator/v10"
)

//UserUsecase ...
type UserUsecase struct {
	userRepo       domain.UserRepo
	contextTimeout time.Duration
}

// NewUserUsecase ...
func NewUserUsecase(a domain.UserRepo, timeout time.Duration) *UserUsecase {
	return &UserUsecase{
		userRepo:       a,
		contextTimeout: timeout,
	}
}

func isRequestValid(m *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Fetch ...
func (u *UserUsecase) Fetch(c context.Context) (users []*domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	users, err = u.userRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return
}

//Store ...
func (u *UserUsecase) Store(c context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	user.HashPassword = user.Password

	if ok, err := isRequestValid(user); !ok {
		return err
	}
	err = u.userRepo.Store(ctx, user)
	if err != nil {
		return
	}
	return
}

//GetByID ...
func (u *UserUsecase) GetByID(c context.Context, id int64) (user *domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	user, err = u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return
}

//Update ...
func (u *UserUsecase) Update(c context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	user.HashPassword = user.Password
	if ok, err := isRequestValid(user); !ok {
		return err
	}
	err = u.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}
	return
}

//DELETE ...
func (u *UserUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	err = u.userRepo.Delete(ctx, id)
	return
}
