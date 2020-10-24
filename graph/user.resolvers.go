package graph

import (
	"context"
	"fmt"

	"github.com/XoronEdge/quizzifire/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := mapInput(input)

	// err := r.uc.Store(ctx, &user)
	// if err != nil {
	// 	return nil, err
	// }
	fmt.Println(user)
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func mapInput(input model.NewUser) (user *model.User) {
	user = &model.User{}
	user.Firstname = input.Firstname
	user.Lastname = input.Lastname
	user.Gender = input.Gender
	user.Email = input.Email
	user.Phone = input.Phone
	user.Password = input.Password
	user.Username = input.Username
	return
}
