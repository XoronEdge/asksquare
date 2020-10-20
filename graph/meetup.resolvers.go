package graph

import (
	"context"
	"fmt"

	"github.com/XoronEdge/quizzifire/graph/model"
)

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	return meetups, nil
}

var meetups = []*model.Meetup{
	{
		ID:          "1",
		Name:        "A meetup in Newyork",
		Description: "Well",
		User: &model.User{
			ID:       "1",
			Username: "lala",
			Email:    "lala@gmail.com",
		},
	},
	{
		ID:          "2",
		Name:        "A meetup in NewJersy",
		Description: "Excellent",
		User: &model.User{
			ID:       "1",
			Username: "lala",
			Email:    "lala@gmail.com",
		},
	},
	{
		ID:          "3",
		Name:        "A meetup in Newhampshire",
		Description: "Amaze",
		User: &model.User{
			ID:       "1",
			Username: "lala",
			Email:    "lala@gmail.com",
		},
	},
}
var users = []model.User{
	{
		ID:       "1",
		Username: "lala",
		Email:    "lala@gmail.com",
	},
}
