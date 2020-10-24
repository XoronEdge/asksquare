package graph

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/XoronEdge/quizzifire/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{}
	fmt.Println(user)
	var ok bool
	result, _ = json.Marshaler(input)
	user = json.Unmarshal(result, user)
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
