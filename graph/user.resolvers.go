package graph

import (
	"context"
	"fmt"

	"github.com/XoronEdge/quizzifire/domain"
	"github.com/XoronEdge/quizzifire/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (guser *model.User, err error) {
	user := mapInput(input)
	err = r.Uc.Store(ctx, user)
	if err != nil {
		fmt.Print("masla hai ", err)
		fmt.Print("*****************", user)
	}

	guser = gmapInput(user)

	return guser, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func mapInput(input model.NewUser) (user *domain.User) {
	user = &domain.User{}
	user.Firstname = *input.Firstname
	user.Lastname = *input.Lastname
	user.Gender = *input.Gender
	user.Email = input.Email
	user.Phone = input.Phone
	user.Password = input.Password
	user.Username = input.Username
	return
}

func gmapInput(input *domain.User) (user *model.User) {
	user = &model.User{}
	user.Firstname = &input.Firstname
	user.Lastname = &input.Lastname
	user.Gender = &input.Gender
	user.Email = input.Email
	user.Phone = input.Phone
	user.Password = input.Password
	user.Username = input.Username
	return
}
