package usecase

import (
	"context"
	"time"

	"github.com/XoronEdge/asksquare/domain"
	validator "github.com/go-playground/validator/v10"
)

type QaReportUsecase struct {
	qrRepo         domain.QaReportRepo
	contextTimeout time.Duration
}

// NewQaReportUsecase ...
func NewQaReportUsecase(a domain.QaReportRepo, timeout time.Duration) *QaReportUsecase {
	return &QaReportUsecase{
		qrRepo:         a,
		contextTimeout: timeout,
	}
}

func isRequestValid(m *domain.QaReport) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Fetch ...
func (qr *QaReportUsecase) Fetch(c context.Context) (qar []*domain.QaReport, err error) {
	ctx, cancel := context.WithTimeout(c, qr.contextTimeout)
	defer cancel()
	qar, err = qr.qrRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return
}

//Store ...
func (qr *QaReportUsecase) Store(c context.Context, qar *domain.QaReport) (err error) {
	ctx, cancel := context.WithTimeout(c, qr.contextTimeout)
	defer cancel()

	qar.UserID = 1
	if ok, err := isRequestValid(qar); !ok {
		return err
	}
	err = qr.qrRepo.Store(ctx, qar)
	if err != nil {
		return
	}
	return
}

//GetByID ...
func (qr *QaReportUsecase) GetByID(c context.Context, id uint) (qar *domain.QaReport, err error) {
	ctx, cancel := context.WithTimeout(c, qr.contextTimeout)
	defer cancel()
	qar, err = qr.qrRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return
}

//Update ...
func (qr *QaReportUsecase) Update(c context.Context, qar *domain.QaReport) (err error) {
	ctx, cancel := context.WithTimeout(c, qr.contextTimeout)
	defer cancel()

	if ok, err := isRequestValid(qar); !ok {
		return err
	}
	qar.UserID = 1
	err = qr.qrRepo.Update(ctx, qar)
	if err != nil {
		return err
	}
	return
}

//Delete ...
func (qr *QaReportUsecase) Delete(c context.Context, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, qr.contextTimeout)
	defer cancel()
	err = qr.qrRepo.Delete(ctx, id)
	return
}

//////Question Hide--------------------------------
type QaHideUsecase struct {
	qhRepo         domain.QaHideRepo
	contextTimeout time.Duration
}

// NewQaHideUsecase ...
func NewQaHideUsecase(a domain.QaHideRepo, timeout time.Duration) *QaHideUsecase {
	return &QaHideUsecase{
		qhRepo:         a,
		contextTimeout: timeout,
	}
}

func isQaHideRequestValid(m *domain.QaHide) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Fetch ...
func (qh *QaHideUsecase) Fetch(c context.Context) (qah []*domain.QaHide, err error) {
	ctx, cancel := context.WithTimeout(c, qh.contextTimeout)
	defer cancel()
	qah, err = qh.qhRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return
}

//Store ...
func (qh *QaHideUsecase) Store(c context.Context, qah *domain.QaHide) (err error) {
	ctx, cancel := context.WithTimeout(c, qh.contextTimeout)
	defer cancel()

	qah.UserID = 1
	if ok, err := isQaHideRequestValid(qah); !ok {
		return err
	}
	err = qh.qhRepo.Store(ctx, qah)
	if err != nil {
		return
	}
	return
}

//GetByID ...
func (qh *QaHideUsecase) GetByID(c context.Context, id uint) (qah *domain.QaHide, err error) {
	ctx, cancel := context.WithTimeout(c, qh.contextTimeout)
	defer cancel()
	qah, err = qh.qhRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return
}

//Update ...
func (qh *QaHideUsecase) Update(c context.Context, qah *domain.QaHide) (err error) {
	ctx, cancel := context.WithTimeout(c, qh.contextTimeout)
	defer cancel()

	if ok, err := isQaHideRequestValid(qah); !ok {
		return err
	}
	qah.UserID = 1
	err = qh.qhRepo.Update(ctx, qah)
	if err != nil {
		return err
	}
	return
}

//Delete ...
func (qh *QaHideUsecase) Delete(c context.Context, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, qh.contextTimeout)
	defer cancel()
	err = qh.qhRepo.Delete(ctx, id)
	return
}

//////Question Answer Later--------------------------------

//QaAnswerLaterUsecase ...
type QaAnswerLaterUsecase struct {
	qalRepo        domain.QaAnswerLaterRepo
	contextTimeout time.Duration
}

// NewQaAnswerLaterUsecase ...
func NewQaAnswerLaterUsecase(a domain.QaAnswerLaterRepo, timeout time.Duration) *QaAnswerLaterUsecase {
	return &QaAnswerLaterUsecase{
		qalRepo:        a,
		contextTimeout: timeout,
	}
}

func isQaAnswerLaterRequestValid(m *domain.QaAnswerLater) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Fetch ...
func (qal *QaAnswerLaterUsecase) Fetch(c context.Context) (qaal []*domain.QaAnswerLater, err error) {
	ctx, cancel := context.WithTimeout(c, qal.contextTimeout)
	defer cancel()
	qaal, err = qal.qalRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return
}

//Store ...
func (qal *QaAnswerLaterUsecase) Store(c context.Context, qaal *domain.QaAnswerLater) (err error) {
	ctx, cancel := context.WithTimeout(c, qal.contextTimeout)
	defer cancel()

	qaal.UserID = 1
	if ok, err := isQaAnswerLaterRequestValid(qaal); !ok {
		return err
	}
	err = qal.qalRepo.Store(ctx, qaal)
	if err != nil {
		return
	}
	return
}

//GetByID ...
func (qal *QaAnswerLaterUsecase) GetByID(c context.Context, id uint) (qaal *domain.QaAnswerLater, err error) {
	ctx, cancel := context.WithTimeout(c, qal.contextTimeout)
	defer cancel()
	qaal, err = qal.qalRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return
}

//Update ...
func (qal *QaAnswerLaterUsecase) Update(c context.Context, qaal *domain.QaAnswerLater) (err error) {
	ctx, cancel := context.WithTimeout(c, qal.contextTimeout)
	defer cancel()

	if ok, err := isQaAnswerLaterRequestValid(qaal); !ok {
		return err
	}
	qaal.UserID = 1
	err = qal.qalRepo.Update(ctx, qaal)
	if err != nil {
		return err
	}
	return
}

//Delete ...
func (qal *QaAnswerLaterUsecase) Delete(c context.Context, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, qal.contextTimeout)
	defer cancel()
	err = qal.qalRepo.Delete(ctx, id)
	return
}
