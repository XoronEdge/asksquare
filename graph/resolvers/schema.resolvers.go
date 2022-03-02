package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"github.com/XoronEdge/asksquare/graph/generated"
)

// QaReport returns generated.QaReportResolver implementation.
func (r *Resolver) QaReport() generated.QaReportResolver { return &qaReportResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
