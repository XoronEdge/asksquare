package graph

import (
	"context"
	"fmt"
	"strconv"

	"github.com/XoronEdge/asksquare/domain"
	"github.com/XoronEdge/asksquare/graph/model"
)

func (r *mutationResolver) CreateQaReport(ctx context.Context, input model.NewQaReport) (*domain.QaReport, error) {
	qar := inpToQRMod(input)
	err := r.QRc.Store(ctx, qar)
	if err != nil {
		return nil, err
	}
	return qar, nil
}

func (r *mutationResolver) UpdateQaReport(ctx context.Context, id string, input model.NewQaReport) (*domain.QaReport, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	qar, err := r.QRc.GetByID(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	qar = upInpToQRMod(input, qar)

	err = r.QRc.Update(ctx, qar)
	if err != nil {
		return nil, err
	}
	return qar, nil
}

func (r *mutationResolver) DeleteQaReport(ctx context.Context, id string) (*domain.QaReport, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QaReports(ctx context.Context, questionID string) ([]*domain.QaReport, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QaReport(ctx context.Context, id string) (*domain.QaReport, error) {
	panic(fmt.Errorf("not implemented"))
}

//ID ...
func (r *qaReportResolver) ID(ctx context.Context, obj *domain.QaReport) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *qaReportResolver) QuestionID(ctx context.Context, obj *domain.QaReport) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *qaReportResolver) User(ctx context.Context, obj *domain.QaReport) (*domain.User, error) {
	panic(fmt.Errorf("not implemented"))
}

type qaReportResolver struct{ *Resolver }

//Input type To Question Report Model
func inpToQRMod(input model.NewQaReport) *domain.QaReport {
	qar := &domain.QaReport{}
	qar.Reason = input.Reason
	qar.Description = input.Description
	qar.QaQuestionID = uint(input.QuestionID)
	return qar
}

//Update Input type To Question Report Model
func upInpToQRMod(input model.NewQaReport, user *domain.QaReport) *domain.QaReport {
	qar := &domain.QaReport{}
	qar.Reason = input.Reason
	qar.Description = input.Description
	qar.QaQuestionID = uint(input.QuestionID)
	return qar
}
