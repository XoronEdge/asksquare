package graph

import (
	"context"
	"strconv"

	"github.com/XoronEdge/asksquare/domain"
	"github.com/XoronEdge/asksquare/graph/generated"
	"github.com/XoronEdge/asksquare/graph/model"
)

func (r *mutationResolver) CreateQaReport(ctx context.Context, input model.NewQaReport) (*domain.QaReport, error) {
	qar := inpToQRMod(input)
	err := r.Di.QRc.Store(ctx, qar)

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
	qar, err := r.Di.QRc.GetByID(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	qar = upInpToQRMod(input, qar)

	err = r.Di.QRc.Update(ctx, qar)
	if err != nil {
		return nil, err
	}
	return qar, nil
}

func (r *mutationResolver) DeleteQaReport(ctx context.Context, id string) (*domain.QaReport, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	qaReport, err := r.Di.QRc.GetByID(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	err = r.Di.QRc.Delete(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	return qaReport, nil
}

func (r *queryResolver) QaReports(ctx context.Context, questionID string) ([]*domain.QaReport, error) {
	_, err := strconv.Atoi(questionID)
	if err != nil {
		return nil, err
	}
	qaReports, err := r.Di.QRc.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return qaReports, nil
}

func (r *queryResolver) QaReport(ctx context.Context, id string) (*domain.QaReport, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	qaReport, err := r.Di.QRc.GetByID(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	return qaReport, nil
}

//ID ...
func (r *qaReportResolver) ID(ctx context.Context, obj *domain.QaReport) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *qaReportResolver) QuestionID(ctx context.Context, obj *domain.QaReport) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *qaReportResolver) User(ctx context.Context, obj *domain.QaReport) (*domain.User, error) {
	user, err := r.Di.Uc.GetByID(ctx, int64(obj.UserID))
	if err != nil {
		return nil, err
	}
	return user, nil
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
func upInpToQRMod(input model.NewQaReport, qar *domain.QaReport) *domain.QaReport {
	qar.Reason = input.Reason
	qar.Description = input.Description
	qar.QaQuestionID = uint(input.QuestionID)
	return qar
}

///Question Hide--------------------------------
func (r *mutationResolver) CreateQaHide(ctx context.Context, input model.NewQaHide) (*domain.QaHide, error) {
	qah := inpToQHMod(input)
	err := r.Di.QHc.Store(ctx, qah)
	if err != nil {
		return nil, err
	}
	return qah, nil
}

func (r *mutationResolver) UpdateQaHide(ctx context.Context, id string, input model.NewQaHide) (*domain.QaHide, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	qah, err := r.Di.QHc.GetByID(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	qah = upInpToQHMod(input, qah)

	err = r.Di.QHc.Update(ctx, qah)
	if err != nil {
		return nil, err
	}
	return qah, nil
}

func (r *mutationResolver) DeleteQaHide(ctx context.Context, id string) (*domain.QaHide, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	qaHide, err := r.Di.QHc.GetByID(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	err = r.Di.QHc.Delete(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	return qaHide, nil
}

func (r *qaHideResolver) ID(ctx context.Context, obj *domain.QaHide) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *qaHideResolver) HideByUser(ctx context.Context, obj *domain.QaHide) (int, error) {
	return int(obj.HideByUser), nil
}

func (r *qaHideResolver) QuestionID(ctx context.Context, obj *domain.QaHide) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *qaHideResolver) User(ctx context.Context, obj *domain.QaHide) (*domain.User, error) {
	user, err := r.Di.Uc.GetByID(ctx, int64(obj.UserID))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) QaHides(ctx context.Context, userID string) ([]*domain.QaHide, error) {
	_, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}
	qaHidess, err := r.Di.QHc.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return qaHidess, nil
}

func (r *queryResolver) QaHide(ctx context.Context, id string) (*domain.QaHide, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	qaHide, err := r.Di.QHc.GetByID(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	return qaHide, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// QaHide returns generated.QaHideResolver implementation.
func (r *Resolver) QaHide() generated.QaHideResolver { return &qaHideResolver{r} }

type qaHideResolver struct{ *Resolver }

//Input type To Question Report Model
func inpToQHMod(input model.NewQaHide) *domain.QaHide {
	qar := &domain.QaHide{}
	qar.HideByUser = uint(input.HideByUser)
	qar.QaQuestionID = uint(input.QuestionID)
	return qar
}

//Update Input type To Question Report Model
func upInpToQHMod(input model.NewQaHide, qar *domain.QaHide) *domain.QaHide {
	qar.HideByUser = uint(input.HideByUser)
	qar.QaQuestionID = uint(input.QuestionID)
	return qar
}

///Question Answer Later --------------------------------
func (r *mutationResolver) CreateQaAnswerLater(ctx context.Context, input model.NewQaAnswerLater) (*domain.QaAnswerLater, error) {
	qaal := inpToQALMod(input)
	err := r.Di.QALc.Store(ctx, qaal)
	if err != nil {
		return nil, err
	}
	return qaal, nil
}

func (r *mutationResolver) UpdateQaAnswerLater(ctx context.Context, id string, input model.NewQaAnswerLater) (*domain.QaAnswerLater, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	qaal, err := r.Di.QALc.GetByID(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	qaal = upInpToQALMod(input, qaal)

	err = r.Di.QALc.Update(ctx, qaal)
	if err != nil {
		return nil, err
	}
	return qaal, nil
}

func (r *mutationResolver) DeleteQaAnswerLater(ctx context.Context, id string) (*domain.QaAnswerLater, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	qaAnswerLater, err := r.Di.QALc.GetByID(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	err = r.Di.QALc.Delete(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	return qaAnswerLater, nil
}

func (r *qaAnswerLaterResolver) ID(ctx context.Context, obj *domain.QaAnswerLater) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *qaAnswerLaterResolver) QuestionID(ctx context.Context, obj *domain.QaAnswerLater) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *qaAnswerLaterResolver) User(ctx context.Context, obj *domain.QaAnswerLater) (*domain.User, error) {
	user, err := r.Di.Uc.GetByID(ctx, int64(obj.UserID))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) QaAnswerLaters(ctx context.Context, userID string) ([]*domain.QaAnswerLater, error) {
	_, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}
	qaAnswerLaterss, err := r.Di.QALc.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return qaAnswerLaterss, nil
}

func (r *queryResolver) QaAnswerLater(ctx context.Context, id string) (*domain.QaAnswerLater, error) {
	value, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	qaAnswerLater, err := r.Di.QALc.GetByID(ctx, uint(value))
	if err != nil {
		return nil, err
	}
	return qaAnswerLater, nil
}

// QaAnswerLater returns generated.QaAnswerLaterResolver implementation.
func (r *Resolver) QaAnswerLater() generated.QaAnswerLaterResolver { return &qaAnswerLaterResolver{r} }

type qaAnswerLaterResolver struct{ *Resolver }

//Input type To Question Report Model
func inpToQALMod(input model.NewQaAnswerLater) *domain.QaAnswerLater {
	qar := &domain.QaAnswerLater{}
	qar.QaQuestionID = uint(input.QuestionID)
	return qar
}

//Update Input type To Question Report Model
func upInpToQALMod(input model.NewQaAnswerLater, qar *domain.QaAnswerLater) *domain.QaAnswerLater {
	qar.QaQuestionID = uint(input.QuestionID)
	return qar
}
