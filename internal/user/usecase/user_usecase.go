package usecase

import (
	"context"
	"time"

	"github.com/XoronEdge/asksquare/domain"
	validator "github.com/go-playground/validator/v10"
)

type userUsercase struct {
	userRepo       domain.UserRepo
	contextTimeout time.Duration
}

// NewUserUsecase ...
func NewUserUsecase(a domain.UserRepo, timeout time.Duration) domain.UserUsecase {
	return &userUsercase{
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
func (u *userUsercase) Fetch(c context.Context) (users []*domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	users, err = u.userRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return
}

//Store ...
func (u *userUsercase) Store(c context.Context, user *domain.User) (err error) {
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
func (u *userUsercase) GetByID(c context.Context, id int64) (user *domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	user, err = u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return
}

//Update ...
func (u *userUsercase) Update(c context.Context, user *domain.User) (err error) {
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
func (u *userUsercase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	err = u.userRepo.Delete(ctx, id)
	return
}
