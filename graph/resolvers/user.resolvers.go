package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/XoronEdge/asksquare/domain"
	"github.com/XoronEdge/asksquare/graph/generated"
	"github.com/XoronEdge/asksquare/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*domain.User, error) {
	user := inpToUserMod(input)
	err := r.Uc.Store(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.NewUser) (*domain.User, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	user, err := r.Uc.GetByID(ctx, int64(value))
	if err != nil {
		return nil, err
	}
	user = upInpToUserMod(input, user)

	err = r.Uc.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*domain.User, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	user, err := r.Uc.GetByID(ctx, int64(value))
	if err != nil {
		return nil, err
	}
	err = r.Uc.Delete(ctx, int64(value))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*domain.User, error) {
	users, err := r.Uc.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*domain.User, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	user, err := r.Uc.GetByID(ctx, int64(value))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userResolver) ID(ctx context.Context, obj *domain.User) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *userResolver) QaReport(ctx context.Context, obj *domain.User) ([]*domain.QaReport, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) QaHide(ctx context.Context, obj *domain.User) ([]*domain.QaHide, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

//  Input Type To User Model
func inpToUserMod(input model.NewUser) *domain.User {
	user := &domain.User{}
	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	user.Gender = *input.Gender
	user.FirstName = *input.Firstname
	user.LastName = *input.Lastname
	user.Phone = input.Phone
	fmt.Println("-------------------------")
	fmt.Printf("%+v\n", user)
	return user
}

// update Input Type To user Model
func upInpToUserMod(input model.NewUser, user *domain.User) *domain.User {
	user.FirstName = *input.Firstname
	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	user.Gender = *input.Gender
	user.FirstName = *input.Firstname
	user.LastName = *input.Lastname
	user.Phone = input.Phone
	return user
}
