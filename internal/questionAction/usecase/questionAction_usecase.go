package usecase

import (
	"context"
	"time"

	"github.com/XoronEdge/asksquare/domain"
	validator "github.com/go-playground/validator/v10"
)

type QaReportUsercase struct {
	qrRepo         domain.QaReportRepo
	contextTimeout time.Duration
}

// NewQaReportUsecase ...
func NewQaReportUsecase(a domain.QaReportRepo, timeout time.Duration) domain.QaReportUsecase {
	return &QaReportUsercase{
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
func (qr *QaReportUsercase) Fetch(c context.Context) (qar []*domain.QaReport, err error) {
	ctx, cancel := context.WithTimeout(c, qr.contextTimeout)
	defer cancel()
	qar, err = qr.qrRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return
}

//Store ...
func (qr *QaReportUsercase) Store(c context.Context, qar *domain.QaReport) (err error) {
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
func (qr *QaReportUsercase) GetByID(c context.Context, id uint) (qar *domain.QaReport, err error) {
	ctx, cancel := context.WithTimeout(c, qr.contextTimeout)
	defer cancel()
	qar, err = qr.qrRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return
}

//Update ...
func (qr *QaReportUsercase) Update(c context.Context, qar *domain.QaReport) (err error) {
	ctx, cancel := context.WithTimeout(c, qr.contextTimeout)
	defer cancel()

	if ok, err := isRequestValid(qar); !ok {
		return err
	}
	err = qr.qrRepo.Update(ctx, qar)
	if err != nil {
		return err
	}
	return
}

//DELETE ...
func (qr *QaReportUsercase) Delete(c context.Context, id uint) (err error) {
	ctx, cancel := context.WithTimeout(c, qr.contextTimeout)
	defer cancel()
	err = qr.qrRepo.Delete(ctx, id)
	return
}
